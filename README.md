# math-controller

If you assign a number to the value of the 「Number」 resource, math-controller solves 「Fizzbuzz」 problems

## Installation
Install base components Use kubectl

```
kubectl apply -f https://github.com/ymktmk/math-controller/releases/download/v1.0.9/math-controller.yaml
```

## Usage
Deploy Kind: Number

```yaml:
apiVersion: math.ymktmk.github.io/v1beta1
kind: Number
metadata:
  name: number-sample
spec:
  value: 10
```

## Reference

* https://www.takutakahashi.dev/lazy-custom-controller-for-kubernetes/

* https://www.wantedly.com/companies/wantedly/post_articles/339271

* https://zoetrope.github.io/kubebuilder-training/introduction/basics.html

* https://cadc.cyberagent.co.jp/2022/program/kubernetes-controller-improvements/

* https://developers.cyberagent.co.jp/blog/archives/36200/

* https://atmarkit.itmedia.co.jp/ait/articles/2109/10/news013.html

