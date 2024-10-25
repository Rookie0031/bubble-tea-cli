

# dev-eks 접근하는 경우

`aws eks update-kubeconfig --region ap-northeast-2 --name ${eks} --role-arn arn:aws:iam::${account_id}}:role/${role_name}}`

### 복사하고 싶은 로컬 경로 지정 후 명령어 실행

```jsx
kubectl cp ${pod_file_path}} ${local_dir}
```
