# gobench2xml

`gobench2xml` is a simple tool that converts go benchmark results to xml. You might find it useful when plotting `go test` results in Jenkins (see [Plot Plugin](https://wiki.jenkins-ci.org/display/JENKINS/Plot+Plugin)). Note that it only takes the first result if there are multiple results with the same name.

### Install

`go get -u github.com/naver/gobench2xml`


### How to Use

`$ go test -run=NONE -bench=. | gobench2xml`

### Example

```sh
$ cat sample.txt

BenchmarkObjectCRUD/POST-4          3000	  13002649 ns/op
BenchmarkObjectCRUD/GET-4           3000	  12771154 ns/op
BenchmarkObjectCRUD/PUT-4           3000	  12738328 ns/op
BenchmarkObjectCRUD/DELETE-4        3000	  13137845 ns/op
PASS

$ cat sample.txt | gobench2xml | tee result.xml
<Benchmarks>
  <NsPerOp>
    <BenchmarkObjectCRUD-PUT>1.2738328e+07</BenchmarkObjectCRUD-PUT>
    <BenchmarkObjectCRUD-DELETE>1.3137845e+07</BenchmarkObjectCRUD-DELETE>
    <BenchmarkObjectCRUD-POST>1.3002649e+07</BenchmarkObjectCRUD-POST>
    <BenchmarkObjectCRUD-GET>1.2771154e+07</BenchmarkObjectCRUD-GET>
  </NsPerOp>
  <AllocedBytesPerOp>
    <BenchmarkObjectCRUD-PUT>0</BenchmarkObjectCRUD-PUT>
    <BenchmarkObjectCRUD-DELETE>0</BenchmarkObjectCRUD-DELETE>
    <BenchmarkObjectCRUD-POST>0</BenchmarkObjectCRUD-POST>
    <BenchmarkObjectCRUD-GET>0</BenchmarkObjectCRUD-GET>
  </AllocedBytesPerOp>
  <AllocsPerOp>
    <BenchmarkObjectCRUD-DELETE>0</BenchmarkObjectCRUD-DELETE>
    <BenchmarkObjectCRUD-POST>0</BenchmarkObjectCRUD-POST>
    <BenchmarkObjectCRUD-GET>0</BenchmarkObjectCRUD-GET>
    <BenchmarkObjectCRUD-PUT>0</BenchmarkObjectCRUD-PUT>
  </AllocsPerOp>
  <MBPerS>
    <BenchmarkObjectCRUD-POST>0</BenchmarkObjectCRUD-POST>
    <BenchmarkObjectCRUD-GET>0</BenchmarkObjectCRUD-GET>
    <BenchmarkObjectCRUD-PUT>0</BenchmarkObjectCRUD-PUT>
    <BenchmarkObjectCRUD-DELETE>0</BenchmarkObjectCRUD-DELETE>
  </MBPerS>
</Benchmarks>
```

### LICENSE

```
Copyright (c) 2017 NAVER Corp.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```
