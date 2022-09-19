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

