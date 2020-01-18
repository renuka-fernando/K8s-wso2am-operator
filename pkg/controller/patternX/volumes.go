package patternX

import (
	apimv1alpha1 "github.com/wso2-incubator/wso2am-k8s-operator/pkg/apis/apim/v1alpha1"
	corev1 "k8s.io/api/core/v1"

)


func getApimXVolumes(apimanager *apimv1alpha1.APIManager, r apimv1alpha1.Profile) ([]corev1.VolumeMount, []corev1.Volume) {

	defaultdeployConf :=  r.Deployment.Configmaps.DeploymentConfigmap
	defaultsynapseconf := "wso2am-p1-am-synapse-configs"
	defaultexecutionconf :=	"wso2am-p1-am-execution-plans"

	var amXvolumemounts []corev1.VolumeMount
	var amXvolume []corev1.Volume

	//adding default deploymentConfigmap
	amXvolumemounts=append(amXvolumemounts,corev1.VolumeMount{
		Name: "wso2am-px-apim-conf",
		MountPath: "/home/wso2carbon/wso2-config-volume/repository/conf/deployment.toml",
		SubPath:"deployment.toml",
	})
	//adding default synapseConfigs pvc
	amXvolumemounts=append(amXvolumemounts,corev1.VolumeMount{
		Name:      "wso2am-px-synapse-conf",
		MountPath: "/home/wso2carbon/wso2-artifact-volume/repository/deployment/server/synapse-configs",

	})
	//adding default executionPlans pvc
	amXvolumemounts=append(amXvolumemounts,corev1.VolumeMount{
		Name:        "wso2am-px-execution-conf",
		MountPath:"/home/wso2carbon/wso2-artifact-volume/repository/deployment/server/executionplans",
	})

	amXvolume =append(amXvolume,corev1.Volume{
		Name: "wso2am-px-apim-conf",
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: defaultdeployConf,
				},
			},
		},
	})
	amXvolume =append(amXvolume,corev1.Volume{
		Name: "wso2am-px-synapse-conf",
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName:defaultsynapseconf,
			},
		},
	})
	amXvolume =append(amXvolume,corev1.Volume{
		Name: "wso2am-px-execution-conf",
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName:defaultexecutionconf,
			},
		},
	})

	return amXvolumemounts, amXvolume

}
