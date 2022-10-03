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

## RBAC

Since it has a cache mechanism, even if only Get is called on implementation, `List` and `Watch` are required as resource access privileges.

## Reference

KubernetesのCRDまわりを整理する。
* https://qiita.com/cvusk/items/773e222e0971a5391a51

Kubernetesに無限の可能性を生み出す「Operator」「CRD」「カスタムコントローラー」とは
* https://atmarkit.itmedia.co.jp/ait/articles/2109/10/news013.html

カスタムコントローラーの基礎
* https://zoetrope.github.io/kubebuilder-training/introduction/basics.html

【Wantedly】Custom Controllerを使ってK8sを学ぶハンズオン
* https://www.youtube.com/watch?v=HPWZRCOk9NE

CustomControllerを手抜きで作る技術
* https://www.takutakahashi.dev/lazy-custom-controller-for-kubernetes/

CustomControllerを作ってデプロイ開始と終了の検知を実現した話
* https://www.wantedly.com/companies/wantedly/post_articles/339271

CyberAgent Kubernetes基盤における運用フローのController化と継続的な改善
* https://cadc.cyberagent.co.jp/2022/program/kubernetes-controller-improvements/

* https://developers.cyberagent.co.jp/blog/archives/36200/

k8sのカスタムリソースで、CronJobの終了を検知してJobを実行する
* https://www.m3tech.blog/entry/k8s-eventjob

Kubernetes Casual Talk 〜Ubie、CA、メルペイのカスタムコントローラー〜
* https://www.youtube.com/watch?v=R2B0cWpIZ5k
