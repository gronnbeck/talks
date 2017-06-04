# High Performance API on Kubernetes

In this talk I will implement a highly scalable read API using go, redis, and Kubernetes. We need a use-case where many more reads than writes will occur. I chose to use a fictional currency index as that example use-case. We'll implement an API and step by step we'll make it scale.

This talk will be presented on NDC 2017.

## Outline

### 1\. The problem

In this section I present the what we're going to present the use-case and the fictional scaling problem.

### 2\. The tools

In this section I present the tools we'll be using to visualize both the currency index and pressure we adding to the service. I'll also give a quick introduction to the performance tool, Redis, and Kubernetes.

### 3\. The Naive Approach

Here I will implement a naive solution to show that the approach will not scale.

### 4\. The Improved Reads Approach

In this section I show that the scaling issue can easily be extended without changes to the API. And also without any changes to the API. And thus no need to plan for these changes.

### 5\. The One With Back-Pressure Control

In this section we'll start testing the write capabilities of the API. I show how this can somehow be mitigate using a queue. But the write API needs to give less guarantees to support the scaling needs.

## Notes

- The performance/load bots in the toolbox was inspired by Kubernetes/contrib [scale-demo](http://blog.kubernetes.io/2015/11/one-million-requests-per-second-dependable-and-dynamic-distributed-systems-at-scale.html)
