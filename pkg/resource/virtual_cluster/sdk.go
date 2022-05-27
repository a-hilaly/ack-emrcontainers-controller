// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package virtual_cluster

import (
	"context"
	"errors"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/emrcontainers"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/emrcontainers-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.EMRContainers{}
	_ = &svcapitypes.VirtualCluster{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribeVirtualClusterOutput
	resp, err = rm.sdkapi.DescribeVirtualClusterWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeVirtualCluster", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.VirtualCluster.Arn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.VirtualCluster.Arn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.VirtualCluster.ContainerProvider != nil {
		f1 := &svcapitypes.ContainerProvider{}
		if resp.VirtualCluster.ContainerProvider.Id != nil {
			f1.ID = resp.VirtualCluster.ContainerProvider.Id
		}
		if resp.VirtualCluster.ContainerProvider.Info != nil {
			f1f1 := &svcapitypes.ContainerInfo{}
			if resp.VirtualCluster.ContainerProvider.Info.EksInfo != nil {
				f1f1f0 := &svcapitypes.EKSInfo{}
				if resp.VirtualCluster.ContainerProvider.Info.EksInfo.Namespace != nil {
					f1f1f0.Namespace = resp.VirtualCluster.ContainerProvider.Info.EksInfo.Namespace
				}
				f1f1.EKSInfo = f1f1f0
			}
			f1.Info = f1f1
		}
		if resp.VirtualCluster.ContainerProvider.Type != nil {
			f1.Type = resp.VirtualCluster.ContainerProvider.Type
		}
		ko.Spec.ContainerProvider = f1
	} else {
		ko.Spec.ContainerProvider = nil
	}
	if resp.VirtualCluster.Id != nil {
		ko.Status.ID = resp.VirtualCluster.Id
	} else {
		ko.Status.ID = nil
	}
	if resp.VirtualCluster.Name != nil {
		ko.Spec.Name = resp.VirtualCluster.Name
	} else {
		ko.Spec.Name = nil
	}
	if resp.VirtualCluster.Tags != nil {
		f6 := map[string]*string{}
		for f6key, f6valiter := range resp.VirtualCluster.Tags {
			var f6val string
			f6val = *f6valiter
			f6[f6key] = &f6val
		}
		ko.Spec.Tags = f6
	} else {
		ko.Spec.Tags = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Status.ID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeVirtualClusterInput, error) {
	res := &svcsdk.DescribeVirtualClusterInput{}

	if r.ko.Status.ID != nil {
		res.SetId(*r.ko.Status.ID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateVirtualClusterOutput
	_ = resp
	resp, err = rm.sdkapi.CreateVirtualClusterWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateVirtualCluster", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.Arn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.Arn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Id != nil {
		ko.Status.ID = resp.Id
	} else {
		ko.Status.ID = nil
	}
	if resp.Name != nil {
		ko.Spec.Name = resp.Name
	} else {
		ko.Spec.Name = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateVirtualClusterInput, error) {
	res := &svcsdk.CreateVirtualClusterInput{}

	if r.ko.Spec.ContainerProvider != nil {
		f0 := &svcsdk.ContainerProvider{}
		if r.ko.Spec.ContainerProvider.ID != nil {
			f0.SetId(*r.ko.Spec.ContainerProvider.ID)
		}
		if r.ko.Spec.ContainerProvider.Info != nil {
			f0f1 := &svcsdk.ContainerInfo{}
			if r.ko.Spec.ContainerProvider.Info.EKSInfo != nil {
				f0f1f0 := &svcsdk.EksInfo{}
				if r.ko.Spec.ContainerProvider.Info.EKSInfo.Namespace != nil {
					f0f1f0.SetNamespace(*r.ko.Spec.ContainerProvider.Info.EKSInfo.Namespace)
				}
				f0f1.SetEksInfo(f0f1f0)
			}
			f0.SetInfo(f0f1)
		}
		if r.ko.Spec.ContainerProvider.Type != nil {
			f0.SetType(*r.ko.Spec.ContainerProvider.Type)
		}
		res.SetContainerProvider(f0)
	}
	if r.ko.Spec.Name != nil {
		res.SetName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.Tags != nil {
		f2 := map[string]*string{}
		for f2key, f2valiter := range r.ko.Spec.Tags {
			var f2val string
			f2val = *f2valiter
			f2[f2key] = &f2val
		}
		res.SetTags(f2)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteVirtualClusterOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteVirtualClusterWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteVirtualCluster", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteVirtualClusterInput, error) {
	res := &svcsdk.DeleteVirtualClusterInput{}

	if r.ko.Status.ID != nil {
		res.SetId(*r.ko.Status.ID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.VirtualCluster,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "ValidationException",
		"ResourceNotFoundException",
		"InternalServerException":
		return true
	default:
		return false
	}
}