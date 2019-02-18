## GoJSON
`GoJSON` is a fast and easy package to marshal/unmarshal struct to/from json. You can use `GoJSON` tool to generate marshal/unmarshal code, in benchmark tests, `GoJSON`'s generate code is almost 6~7 times faster than `encoding/json`.

## Example 
```golang
# install
  go get -u -v github.com/go-fish/gojson
  cd ${GOPATH}/src/github.com/go-fish/gojson/cmd
  go build -o gojson main.go


# usage
  gojson [options] <input dir|file>
  
  -inline
        Use inline function in generate code (default true)
  -m string
        Mode of generate, eg: encode, decode, all (default "all")
  -o string
        Optional name of the output file to be generated. (default "gojson.generate.go")
  -unsafe
        Use decoder without copy data
  -version
        Show version information.
        
```
 

with -inline && -unsafe, you can get most performance generate code
with -m, gojson will generate marshal/unmarshal/both for all expose structs in input dir/file.

For expose structs, gojson generate `MarshalJSON/UnmarshalJSON` methods for marshal/unmarshal json. You also can use `gojson.Marshal/gojson.Unmarshal` functions to marshal/unmarshal json.

## Benchmark
### Large Payload
#### Unmarshal
<table>
	<tr>
		<th>gojson</th>
		<th>114658 ns/op</th>
		<th>457.72 MB/s</th>
		<th>5873 B/op</th>
		<th>253 allocs/op</th>
	</tr>
	<tr>
		<th>jsonparser</th>
		<th>116647 ns/op</th>
		<th>449.92 MB/s</th>
		<th>0 B/op</th>
		<th>0 allocs/op</th>
	</tr>
	<tr>
		<th>easyjson</th>
		<th>219673 ns/op</th>
		<th>238.91 MB/s</th>
		<th>9264 B/op</th>
		<th>349 allocs/op</th>
	</tr>
	<tr>
		<th>gojay</th>
		<th>92398 ns/op</th>
		<th>567.99 MB/s</th>
		<th>5985 B/op</th>
		<th>259 allocs/op</th>
	</tr>
	<tr>
		<th>jsoniter</th>
		<th>159877 ns/op</th>
		<th>328.26 MB/s</th>
		<th>800 B/op</th>
		<th>29 allocs/op</th>
	</tr>
	<tr>
		<th>encoding/json</th>
		<th>680867 ns/op</th>
		<th>377.08 MB/s</th>
		<th>6154 B/op</th>
		<th>258 allocs/op</th>
	</tr>
</table>

#### Marshal
<table>
	<tr>
		<th>gojson</th>
		<th>131 ns/op</th>
		<th>400511.56 MB/s</th>
		<th>32 B/op</th>
		<th>1 allocs/op</th>
	</tr>
	<tr>
		<th>easyjson</th>
		<th>235 ns/op</th>
		<th>223039.53 MB/s</th>
		<th>272 B/op</th>
		<th>3 allocs/op</th>
	</tr>
	<tr>
		<th>encoding/json</th>
		<th>597 ns/op</th>
		<th>87819.16 MB/s</th>
		<th>120 B/op</th>
		<th>4 allocs/op</th>
	</tr>
</table>

### Medium Payload
#### Unmarshal

<table>
	<tr>
		<th>gojson</th>
		<th>4056 ns/op</th>
		<th>510.51 MB/s</th>
		<th>240 B/op</th>
		<th>10 allocs/op</th>
	</tr>
	<tr>
		<th>jsonparser</th>
		<th>9379 ns/op</th>
		<th>220.81 MB/s</th>
		<th>0 B/op</th>
		<th>0 allocs/op</th>
	</tr>
	<tr>
		<th>easyjson</th>
		<th>8107 ns/op</th>
		<th>255.43 MB/s</th>
		<th>288 B/op</th>
		<th>11 allocs/op</th>
	</tr>
	<tr>
		<th>gojay</th>
		<th>5004 ns/op</th>
		<th>413.84 MB/s</th>
		<th>2472 B/op</th>
		<th>10 allocs/op</th>
	</tr>
	<tr>
		<th>jsoniter</th>
		<th>10579 ns/op</th>
		<th>195.75 MB/s</th>
		<th>2576 B/op</th>
		<th>12 allocs/op</th>
	</tr>
	<tr>
		<th>encoding/json</th>
		<th>28030 ns/op</th>
		<th>73.88 MB/s</th>
		<th>520 B/op</th>
		<th>15 allocs/op</th>
	</tr>
</table>

#### Marshal
<table>
	<tr>
		<th>gojson</th>
		<th>105 ns/op</th>
		<th>19551.00 MB/s</th>
		<th>32 B/op</th>
		<th>1 allocs/op</th>
	</tr>
	<tr>
		<th>easyjson</th>
		<th>193 ns/op</th>
		<th>10700.99 MB/s</th>
		<th>272 B/op</th>
		<th>3 allocs/op</th>
	</tr>
	<tr>
		<th>encoding/json</th>
		<th>443 ns/op</th>
		<th>4667.11 MB/s</th>
		<th>120 B/op</th>
		<th>4 allocs/op</th>
	</tr>
</table>

### Small Payload
#### Unmarshal

<table>
	<tr>
		<th>gojson</th>
		<th>812 ns/op</th>
		<th>176.00 MB/s</th>
		<th>160 B/op</th>
		<th>5 allocs/op</th>
	</tr>
	<tr>
		<th>jsonparser</th>
		<th>925 ns/op</th>
		<th>154.47 MB/s</th>
		<th>0 B/op</th>
		<th>0 allocs/op</th>
	</tr>
	<tr>
		<th>easyjson</th>
		<th>1209 ns/op</th>
		<th>118.26 MB/s</th>
		<th>128 B/op</th>
		<th>8 allocs/op</th>
	</tr>
	<tr>
		<th>gojay</th>
		<th>977 ns/op</th>
		<th>146.30 MB/s</th>
		<th>304 B/op</th>
		<th>6 allocs/op</th>
	</tr>
	<tr>
		<th>jsoniter</th>
		<th>1847 ns/op</th>
		<th>77.42 MB/s</th>
		<th>336 B/op</th>
		<th>7 allocs/op</th>
	</tr>
	<tr>
		<th>encoding/json</th>
		<th>2576 ns/op</th>
		<th>55.49 MB/s</th>
		<th>328 B/op</th>
		<th>7 allocs/op</th>
	</tr>
</table>

#### Marshal
<table>
	<tr>
		<th>gojson</th>
		<th>331 ns/op</th>
		<th>431.90 MB/s</th>
		<th>32 B/op</th>
		<th>1 allocs/op</th>
	</tr>
	<tr>
		<th>easyjson</th>
		<th>439 ns/op</th>
		<th>325.35 MB/s</th>
		<th>272 B/op</th>
		<th>3 allocs/op</th>
	</tr>
	<tr>
		<th>encoding/json</th>
		<th>1398 ns/op</th>
		<th>102.22 MB/s</th>
		<th>120 B/op</th>
		<th>4 allocs/op</th>
	</tr>
</table>

## Questions
Any questions or bugs can go though github issues.