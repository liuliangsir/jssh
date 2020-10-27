package jsbuiltin

var modules []string

func init() {
	// cli.js
	modules = append(modules, "Y29uc3QgY2xpID0ge307Cgp7CiAgY29uc3QgYXJncyA9IFtdOwogIGNvbnN0IG9wdHMgPSB7fTsKCiAgZnVuY3Rpb24gZ2V0RmxhZ05hbWUocykgewogICAgaWYgKHMuc3RhcnRzV2l0aCgiLS0iKSkgewogICAgICByZXR1cm4gcy5zbGljZSgyKTsKICAgIH0KICAgIGlmIChzLnN0YXJ0c1dpdGgoIi0iKSkgewogICAgICByZXR1cm4gcy5zbGljZSgxKTsKICAgIH0KICB9CgogIGZvciAobGV0IGkgPSAyOyBpIDwgX19hcmdzLmxlbmd0aDsgaSsrKSB7CiAgICBjb25zdCB2ID0gX19hcmdzW2ldOwogICAgY29uc3QgdjIgPSBfX2FyZ3NbaSArIDFdOwogICAgaWYgKHYuc3RhcnRzV2l0aCgiLSIpKSB7CiAgICAgIGNvbnN0IHIgPSB2Lm1hdGNoKC9eLS0/KFtcd1wtX10rKT0oLiopJC8pOwogICAgICBpZiAocikgewogICAgICAgIG9wdHNbclsxXV0gPSByWzJdOwogICAgICB9IGVsc2UgewogICAgICAgIGlmICh2MiAhPT0gdW5kZWZpbmVkKSB7CiAgICAgICAgICBpZiAodjIuc3RhcnRzV2l0aCgiLSIpKSB7CiAgICAgICAgICAgIG9wdHNbZ2V0RmxhZ05hbWUodildID0gdHJ1ZTsKICAgICAgICAgIH0gZWxzZSB7CiAgICAgICAgICAgIG9wdHNbZ2V0RmxhZ05hbWUodildID0gdjI7CiAgICAgICAgICAgIGkrKzsKICAgICAgICAgIH0KICAgICAgICB9IGVsc2UgewogICAgICAgICAgb3B0c1tnZXRGbGFnTmFtZSh2KV0gPSB0cnVlOwogICAgICAgIH0KICAgICAgfQogICAgfSBlbHNlIHsKICAgICAgYXJncy5wdXNoKHYpOwogICAgfQogIH0KCiAgY2xpLmdldCA9IGZ1bmN0aW9uIChuKSB7CiAgICBpZiAodHlwZW9mIG4gPT09ICJudW1iZXIiKSB7CiAgICAgIHJldHVybiBhcmdzW25dOwogICAgfSBlbHNlIHsKICAgICAgcmV0dXJuIG9wdHNbbl07CiAgICB9CiAgfTsKCiAgY2xpLmJvb2wgPSBmdW5jdGlvbiAobikgewogICAgaWYgKG9wdHNbbl0gPT09IGZhbHNlIHx8IG9wdHNbbl0gPT09IHVuZGVmaW5lZCkgcmV0dXJuIGZhbHNlOwogICAgaWYgKG9wdHNbbl0gPT09IHRydWUpIHJldHVybiB0cnVlOwogICAgY29uc3QgcyA9IG9wdHNbbl0udG9Mb3dlckNhc2UoKTsKICAgIHJldHVybiAhKHMgPT09ICIwIiB8fCBzID09PSAiZiIgfHwgcyA9PT0gImZhbHNlIik7CiAgfTsKCiAgY2xpLmFyZ3MgPSBmdW5jdGlvbiAoKSB7CiAgICByZXR1cm4gWy4uLmFyZ3NdOwogIH07CgogIGNsaS5vcHRzID0gZnVuY3Rpb24gKCkgewogICAgcmV0dXJuIHsgLi4ub3B0cyB9OwogIH07CgogIGNsaS5wcm9tcHQgPSBmdW5jdGlvbiAobWVzc2FnZSkgewogICAgaWYgKG1lc3NhZ2UpIHByaW50KG1lc3NhZ2UpOwogICAgcmV0dXJuIHJlYWRsaW5lKCk7CiAgfTsKCiAgY2xpLl9zdWJjb21tYW5kID0ge307CgogIGNsaS5zdWJjb21tYW5kID0gZnVuY3Rpb24gKG5hbWUsIGNhbGxiYWNrKSB7CiAgICBpZiAodHlwZW9mIGNhbGxiYWNrICE9PSBgZnVuY3Rpb25gKSB7CiAgICAgIHRocm93IG5ldyBUeXBlRXJyb3IoYGNhbGxiYWNrIGV4cGVjdGVkIGEgZnVuY3Rpb25gKTsKICAgIH0KICAgIGlmIChjbGkuX3N1YmNvbW1hbmRbbmFtZV0pIHsKICAgICAgdGhyb3cgbmV3IEVycm9yKGBzdWJjb21tYW5kICR7bmFtZX0gaXMgYWxyZWFkeSByZWdpc3RlcmVkYCk7CiAgICB9CiAgICBjbGkuX3N1YmNvbW1hbmRbbmFtZV0gPSBjYWxsYmFjazsKICB9OwoKICBjbGkuc3ViY29tbWFuZHN0YXJ0ID0gZnVuY3Rpb24gKCkgewogICAgY29uc3QgbmFtZSA9IGNsaS5nZXQoMCk7CiAgICBpZiAoY2xpLl9zdWJjb21tYW5kW25hbWVdKSB7CiAgICAgIHJldHVybiBjbGkuX3N1YmNvbW1hbmRbbmFtZV0oKTsKICAgIH0KICAgIGlmIChjbGkuX3N1YmNvbW1hbmRbYCpgXSkgewogICAgICByZXR1cm4gY2xpLl9zdWJjb21tYW5kW2AqYF0oKTsKICAgIH0KICAgIHRocm93IG5ldyBFcnJvcihgdW5yZWNvZ25pemVkIHN1YmNvbW1hbmQgJHtuYW1lfWApOwogIH07Cn0K")

	// log.js
	modules = append(modules, "ewogIGxvZy5mYXRhbCA9IGZ1bmN0aW9uIChtZXNzYWdlLCAuLi5hcmdzKSB7CiAgICBsb2cuZXJyb3IobWVzc2FnZSwgLi4uYXJncyk7CiAgICBleGl0KDEpOwogIH07Cn0K")
}