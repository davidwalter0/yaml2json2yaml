Description
===================
[![build status](http://k8s-gitlab-ui/ci/projects/3/status.png?ref=master)](http://k8s-gitlab-ui/ci/projects/3?ref=master) 
yaml2json2yaml using transform library based on github.com/lbronze1man/yaml2json

bin/json2yaml,bin/yaml2json built for x86_64 linux

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
