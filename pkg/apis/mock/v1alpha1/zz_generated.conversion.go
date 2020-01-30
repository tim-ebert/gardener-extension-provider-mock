// +build !ignore_autogenerated

/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	mock "github.com/gardener/gardener-extension-provider-mock/pkg/apis/mock"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*CloudControllerManagerConfig)(nil), (*mock.CloudControllerManagerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudControllerManagerConfig_To_mock_CloudControllerManagerConfig(a.(*CloudControllerManagerConfig), b.(*mock.CloudControllerManagerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*mock.CloudControllerManagerConfig)(nil), (*CloudControllerManagerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_mock_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(a.(*mock.CloudControllerManagerConfig), b.(*CloudControllerManagerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CloudProfileConfig)(nil), (*mock.CloudProfileConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudProfileConfig_To_mock_CloudProfileConfig(a.(*CloudProfileConfig), b.(*mock.CloudProfileConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*mock.CloudProfileConfig)(nil), (*CloudProfileConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_mock_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(a.(*mock.CloudProfileConfig), b.(*CloudProfileConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControlPlaneConfig)(nil), (*mock.ControlPlaneConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControlPlaneConfig_To_mock_ControlPlaneConfig(a.(*ControlPlaneConfig), b.(*mock.ControlPlaneConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*mock.ControlPlaneConfig)(nil), (*ControlPlaneConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_mock_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(a.(*mock.ControlPlaneConfig), b.(*ControlPlaneConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InfrastructureConfig)(nil), (*mock.InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureConfig_To_mock_InfrastructureConfig(a.(*InfrastructureConfig), b.(*mock.InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*mock.InfrastructureConfig)(nil), (*InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_mock_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(a.(*mock.InfrastructureConfig), b.(*InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InfrastructureStatus)(nil), (*mock.InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureStatus_To_mock_InfrastructureStatus(a.(*InfrastructureStatus), b.(*mock.InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*mock.InfrastructureStatus)(nil), (*InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_mock_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(a.(*mock.InfrastructureStatus), b.(*InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*WorkerConfig)(nil), (*mock.WorkerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_WorkerConfig_To_mock_WorkerConfig(a.(*WorkerConfig), b.(*mock.WorkerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*mock.WorkerConfig)(nil), (*WorkerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_mock_WorkerConfig_To_v1alpha1_WorkerConfig(a.(*mock.WorkerConfig), b.(*WorkerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*WorkerStatus)(nil), (*mock.WorkerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_WorkerStatus_To_mock_WorkerStatus(a.(*WorkerStatus), b.(*mock.WorkerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*mock.WorkerStatus)(nil), (*WorkerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_mock_WorkerStatus_To_v1alpha1_WorkerStatus(a.(*mock.WorkerStatus), b.(*WorkerStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_CloudControllerManagerConfig_To_mock_CloudControllerManagerConfig(in *CloudControllerManagerConfig, out *mock.CloudControllerManagerConfig, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_v1alpha1_CloudControllerManagerConfig_To_mock_CloudControllerManagerConfig is an autogenerated conversion function.
func Convert_v1alpha1_CloudControllerManagerConfig_To_mock_CloudControllerManagerConfig(in *CloudControllerManagerConfig, out *mock.CloudControllerManagerConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudControllerManagerConfig_To_mock_CloudControllerManagerConfig(in, out, s)
}

func autoConvert_mock_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in *mock.CloudControllerManagerConfig, out *CloudControllerManagerConfig, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_mock_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig is an autogenerated conversion function.
func Convert_mock_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in *mock.CloudControllerManagerConfig, out *CloudControllerManagerConfig, s conversion.Scope) error {
	return autoConvert_mock_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in, out, s)
}

func autoConvert_v1alpha1_CloudProfileConfig_To_mock_CloudProfileConfig(in *CloudProfileConfig, out *mock.CloudProfileConfig, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_CloudProfileConfig_To_mock_CloudProfileConfig is an autogenerated conversion function.
func Convert_v1alpha1_CloudProfileConfig_To_mock_CloudProfileConfig(in *CloudProfileConfig, out *mock.CloudProfileConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudProfileConfig_To_mock_CloudProfileConfig(in, out, s)
}

func autoConvert_mock_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(in *mock.CloudProfileConfig, out *CloudProfileConfig, s conversion.Scope) error {
	return nil
}

// Convert_mock_CloudProfileConfig_To_v1alpha1_CloudProfileConfig is an autogenerated conversion function.
func Convert_mock_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(in *mock.CloudProfileConfig, out *CloudProfileConfig, s conversion.Scope) error {
	return autoConvert_mock_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(in, out, s)
}

func autoConvert_v1alpha1_ControlPlaneConfig_To_mock_ControlPlaneConfig(in *ControlPlaneConfig, out *mock.ControlPlaneConfig, s conversion.Scope) error {
	out.CloudControllerManager = (*mock.CloudControllerManagerConfig)(unsafe.Pointer(in.CloudControllerManager))
	return nil
}

// Convert_v1alpha1_ControlPlaneConfig_To_mock_ControlPlaneConfig is an autogenerated conversion function.
func Convert_v1alpha1_ControlPlaneConfig_To_mock_ControlPlaneConfig(in *ControlPlaneConfig, out *mock.ControlPlaneConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControlPlaneConfig_To_mock_ControlPlaneConfig(in, out, s)
}

func autoConvert_mock_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in *mock.ControlPlaneConfig, out *ControlPlaneConfig, s conversion.Scope) error {
	out.CloudControllerManager = (*CloudControllerManagerConfig)(unsafe.Pointer(in.CloudControllerManager))
	return nil
}

// Convert_mock_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig is an autogenerated conversion function.
func Convert_mock_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in *mock.ControlPlaneConfig, out *ControlPlaneConfig, s conversion.Scope) error {
	return autoConvert_mock_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in, out, s)
}

func autoConvert_v1alpha1_InfrastructureConfig_To_mock_InfrastructureConfig(in *InfrastructureConfig, out *mock.InfrastructureConfig, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_InfrastructureConfig_To_mock_InfrastructureConfig is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureConfig_To_mock_InfrastructureConfig(in *InfrastructureConfig, out *mock.InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureConfig_To_mock_InfrastructureConfig(in, out, s)
}

func autoConvert_mock_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *mock.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	return nil
}

// Convert_mock_InfrastructureConfig_To_v1alpha1_InfrastructureConfig is an autogenerated conversion function.
func Convert_mock_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *mock.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_mock_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in, out, s)
}

func autoConvert_v1alpha1_InfrastructureStatus_To_mock_InfrastructureStatus(in *InfrastructureStatus, out *mock.InfrastructureStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_InfrastructureStatus_To_mock_InfrastructureStatus is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureStatus_To_mock_InfrastructureStatus(in *InfrastructureStatus, out *mock.InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureStatus_To_mock_InfrastructureStatus(in, out, s)
}

func autoConvert_mock_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *mock.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	return nil
}

// Convert_mock_InfrastructureStatus_To_v1alpha1_InfrastructureStatus is an autogenerated conversion function.
func Convert_mock_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *mock.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_mock_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in, out, s)
}

func autoConvert_v1alpha1_WorkerConfig_To_mock_WorkerConfig(in *WorkerConfig, out *mock.WorkerConfig, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_WorkerConfig_To_mock_WorkerConfig is an autogenerated conversion function.
func Convert_v1alpha1_WorkerConfig_To_mock_WorkerConfig(in *WorkerConfig, out *mock.WorkerConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_WorkerConfig_To_mock_WorkerConfig(in, out, s)
}

func autoConvert_mock_WorkerConfig_To_v1alpha1_WorkerConfig(in *mock.WorkerConfig, out *WorkerConfig, s conversion.Scope) error {
	return nil
}

// Convert_mock_WorkerConfig_To_v1alpha1_WorkerConfig is an autogenerated conversion function.
func Convert_mock_WorkerConfig_To_v1alpha1_WorkerConfig(in *mock.WorkerConfig, out *WorkerConfig, s conversion.Scope) error {
	return autoConvert_mock_WorkerConfig_To_v1alpha1_WorkerConfig(in, out, s)
}

func autoConvert_v1alpha1_WorkerStatus_To_mock_WorkerStatus(in *WorkerStatus, out *mock.WorkerStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_WorkerStatus_To_mock_WorkerStatus is an autogenerated conversion function.
func Convert_v1alpha1_WorkerStatus_To_mock_WorkerStatus(in *WorkerStatus, out *mock.WorkerStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_WorkerStatus_To_mock_WorkerStatus(in, out, s)
}

func autoConvert_mock_WorkerStatus_To_v1alpha1_WorkerStatus(in *mock.WorkerStatus, out *WorkerStatus, s conversion.Scope) error {
	return nil
}

// Convert_mock_WorkerStatus_To_v1alpha1_WorkerStatus is an autogenerated conversion function.
func Convert_mock_WorkerStatus_To_v1alpha1_WorkerStatus(in *mock.WorkerStatus, out *WorkerStatus, s conversion.Scope) error {
	return autoConvert_mock_WorkerStatus_To_v1alpha1_WorkerStatus(in, out, s)
}
