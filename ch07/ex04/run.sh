#!/bin/bash

cat << "EOF" | go run ch07-ex04.go simplereader.go
<html><head></head>
<body>
  <ul>
      <li><a href="https://example.com/foo">foo</a></li>
      <li><a href="https://example.com/bar">bar</a></li>
      <li><a href="https://example.com/baz">baz</a></li>
  </ul>
  <img src="asdf.png" />
</body>
</html>
EOF
