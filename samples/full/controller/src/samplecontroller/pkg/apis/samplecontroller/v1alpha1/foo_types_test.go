
/*
Copyright 2017 The Kubernetes Authors.

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


package v1alpha1_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

    . "samplecontroller/pkg/apis/samplecontroller/v1alpha1"
    . "samplecontroller/pkg/client/clientset/versioned/typed/samplecontroller/v1alpha1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the Foo resource tests

var _ = Describe("Foo", func() {
    var instance Foo
    var expected Foo
    var client FooInterface

    BeforeEach(func() {
        instance = Foo{}
        instance.Name = "instance-1"

        expected = instance
    })

    AfterEach(func() {
        client.Delete(instance.Name, &metav1.DeleteOptions{})
    })

    // INSERT YOUR CODE HERE - add more "Describe" tests

    // Automatically created storage tests
    Describe("when sending a storage request", func() {
        Context("for a valid config", func() {
            It("should provide CRUD access to the object", func() {
                client = cs.SamplecontrollerV1alpha1().Foos("default")

                By("returning success from the create request")
                actual, err := client.Create(&instance)
                Expect(err).ShouldNot(HaveOccurred())

                By("defaulting the expected fields")
                Expect(actual.Spec).To(Equal(expected.Spec))

                By("returning the item for list requests")
                result, err := client.List(metav1.ListOptions{})
                Expect(err).ShouldNot(HaveOccurred())
                Expect(result.Items).To(HaveLen(1))
                Expect(result.Items[0].Spec).To(Equal(expected.Spec))

                By("returning the item for get requests")
                actual, err = client.Get(instance.Name, metav1.GetOptions{})
                Expect(err).ShouldNot(HaveOccurred())
                Expect(actual.Spec).To(Equal(expected.Spec))

                By("deleting the item for delete requests")
                err = client.Delete(instance.Name, &metav1.DeleteOptions{})
                Expect(err).ShouldNot(HaveOccurred())
                result, err = client.List(metav1.ListOptions{})
                Expect(err).ShouldNot(HaveOccurred())
                Expect(result.Items).To(HaveLen(0))
            })
        })
    })
})
