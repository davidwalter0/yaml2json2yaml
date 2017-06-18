Description
===================
[![Build Status](https://travis-ci.org/davidwalter0/yaml2json2yaml.svg?branch=master)](https://travis-ci.org/davidwalter0/yaml2json2yaml)

yaml2json2yaml using transform library based on github.com/lbronze1man/yaml2json

bin/json2yaml,bin/yaml2json built for x86_64 linux

- Example copy tools from docker container

```docker run --rm -v ${PWD}/target:/target yaml2json2yaml:latest cp /bin/yaml2json /bin/json2yaml /bin/slack-status /target```


Usage
====================
### shell
* find the build of you platform
* run `echo "a: 1" | yaml2json` to see result

### read from file save to file
```
echo "a: 1" | bin/yaml2json > 2.json
```

```
(
echo "a: 1" | bin/yaml2json | tee 1.json | bin/json2yaml | tee 1.yaml | bin/yaml2json | tee 2.json; echo;
cat 1.json; echo ;
cat 1.yaml; echo;
cat 2.json; echo ;
)

{
  "a": 1
}
a: 1

{
  "a": 1
}


```

Reference
====================
removed? https://github.com/peter-edge/yaml2json
https://github.com/lbronze1man/yaml2json
