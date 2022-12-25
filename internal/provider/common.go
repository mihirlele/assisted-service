package provider

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/swag"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/featuresupport"
	"github.com/openshift/assisted-service/models"
	"github.com/pkg/errors"
)

func isPlatformBM(platform *models.Platform) bool {
	return platform != nil && *platform.Type == models.PlatformTypeBaremetal
}

func isPlatformNone(platform *models.Platform) bool {
	return platform != nil && *platform.Type == models.PlatformTypeNone
}

func isClusterPlatformNone(cluster *common.Cluster) bool {
	return cluster != nil && isPlatformNone(cluster.Platform)
}

func isClusterPlatformBM(cluster *common.Cluster) bool {
	return cluster != nil && isPlatformBM(cluster.Platform)
}

func checkPlatformWrongParamsInput(platform *models.Platform, userManagedNetworking *bool, cluster *common.Cluster) error {
	if platform != nil && userManagedNetworking != nil {
		userManagedNetworkingStatus := "enabled"
		if !*userManagedNetworking {
			userManagedNetworkingStatus = "disabled"
		}

		if (!*userManagedNetworking && isPlatformNone(platform)) || (*userManagedNetworking && isPlatformBM(platform)) {
			return common.NewApiError(http.StatusBadRequest, errors.Errorf("Can't set %s platform with user-managed-networking %s", *platform.Type, userManagedNetworkingStatus))
		}
	}

	// If current cluster platform is different than baremetal/none, and we want to set the cluster platform to one
	// of those platforms, that might cause the cluster to be in wrong state (baremetal + umn enabled, none + umn disabled)
	// In those cases return bed request
	if cluster != nil && !isClusterPlatformBM(cluster) && !isClusterPlatformNone(cluster) {
		if platform != nil && isPlatformBM(platform) && swag.BoolValue(cluster.UserManagedNetworking) {
			return common.NewApiError(http.StatusBadRequest, errors.Errorf("Can't set %s platform with user-managed-networking enabled", *platform.Type))
		}

		if platform != nil && isPlatformNone(platform) && !swag.BoolValue(cluster.UserManagedNetworking) {
			return common.NewApiError(http.StatusBadRequest, errors.Errorf("Can't set %s platform with user-managed-networking disabled", *platform.Type))
		}
	}

	return nil
}

func doesPlatformDifferentThanBaremetalOrNone(platform *models.Platform, cluster *common.Cluster) bool {
	if platform != nil && !isPlatformBM(platform) && !isPlatformNone(platform) {
		return true
	}

	if platform == nil && !isClusterPlatformBM(cluster) && !isClusterPlatformNone(cluster) {
		return true
	}

	return false
}

func GetActualUpdateClusterPlatformParams(platform *models.Platform, userManagedNetworking *bool, cluster *common.Cluster) (*models.Platform, *bool, error) {
	if platform == nil && userManagedNetworking == nil {
		return nil, nil, nil
	}

	if err := checkPlatformWrongParamsInput(platform, userManagedNetworking, cluster); err != nil {
		return nil, nil, err
	}

	if doesPlatformDifferentThanBaremetalOrNone(platform, cluster) {
		return platform, userManagedNetworking, nil
	}

	if isClusterPlatformBM(cluster) {
		if !swag.BoolValue(userManagedNetworking) && (platform == nil || isPlatformBM(platform)) {
			// Platform is already baremetal, nothing to do
			return nil, nil, nil
		}

		if (platform != nil && isPlatformNone(platform)) || (swag.BoolValue(userManagedNetworking) && platform == nil) {
			return &models.Platform{Type: common.PlatformTypePtr(models.PlatformTypeNone)}, swag.Bool(true), nil
		}
	} else if isClusterPlatformNone(cluster) {
		if (userManagedNetworking == nil || *userManagedNetworking) && (platform == nil || isPlatformNone(platform)) {
			// Platform is already none, nothing to do
			return nil, nil, nil
		}

		if *cluster.HighAvailabilityMode == models.ClusterHighAvailabilityModeNone {
			if !swag.BoolValue(userManagedNetworking) || (platform != nil && !isPlatformNone(platform)) {
				return nil, nil, common.NewApiError(http.StatusBadRequest, errors.New("disabling User Managed Networking or setting platform different than none platform is not allowed in single node Openshift"))
			}
		}

		if !swag.BoolValue(userManagedNetworking) {
			if platform == nil || isPlatformBM(platform) {
				if cluster.CPUArchitecture != common.X86CPUArchitecture && !featuresupport.IsFeatureSupported(cluster.OpenshiftVersion, models.FeatureSupportLevelFeaturesItems0FeatureIDARM64ARCHITECTUREWITHCLUSTERMANAGEDNETWORKING) {
					return nil, nil, common.NewApiError(http.StatusBadRequest, errors.New("disabling User Managed Networking or setting Bare-Metal platform is not allowed for clusters with non-x86_64 CPU architecture"))
				}

				return &models.Platform{Type: common.PlatformTypePtr(models.PlatformTypeBaremetal)}, swag.Bool(false), nil
			}
		}
	}

	return platform, userManagedNetworking, nil
}

func GetActualCreateClusterPlatformParams(platform *models.Platform, userManagedNetworking *bool, highAvailabilityMode *string) (*models.Platform, *bool, error) {
	if err := checkPlatformWrongParamsInput(platform, userManagedNetworking, nil); err != nil {
		return nil, nil, err
	}

	if platform != nil && !isPlatformBM(platform) && !isPlatformNone(platform) {
		return platform, userManagedNetworking, nil
	}

	if *highAvailabilityMode == models.ClusterHighAvailabilityModeFull {
		if (platform == nil || isPlatformBM(platform)) && !swag.BoolValue(userManagedNetworking) {
			return &models.Platform{Type: common.PlatformTypePtr(models.PlatformTypeBaremetal)}, swag.Bool(false), nil
		}

		if swag.BoolValue(userManagedNetworking) || isPlatformNone(platform) {
			return &models.Platform{Type: common.PlatformTypePtr(models.PlatformTypeNone)}, swag.Bool(true), nil
		}
	} else { // *highAvailabilityMode == models.ClusterHighAvailabilityModeNone
		if isPlatformBM(platform) {
			return nil, nil, common.NewApiError(http.StatusBadRequest, errors.Errorf("Can't set %s platform on single node OpenShift", *platform.Type))
		}

		if userManagedNetworking != nil && !*userManagedNetworking {
			return nil, nil, common.NewApiError(http.StatusBadRequest, errors.New("Can't disable user-managed-networking on single node OpenShift"))
		}

		return &models.Platform{Type: common.PlatformTypePtr(models.PlatformTypeNone)}, swag.Bool(true), nil
	}

	return nil, nil, common.NewApiError(http.StatusBadRequest, errors.Errorf("Got invalid platform (%s) and/or user-managed-networking (%v)", *platform.Type, userManagedNetworking))
}

func PlatformTypeString(ptype *models.PlatformType) string {
	if ptype == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PlatformType: %s", string(*ptype))
}

func PlatformString(platform *models.Platform) string {
	if platform == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Platform: %s", PlatformTypeString(platform.Type))
}
