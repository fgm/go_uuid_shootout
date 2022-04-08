# A comparison of UUID and other unique values generators in Go
## Interpreting the results

These generators do NOT emphasize the same properties, so do not compare apples to oranges.

## Running the shootout

```
go test -bench=. -benchmem -benchtime=1s .
```

## Results

These results are from a 10 runs comparison with
[golang.org/x/perf/cmd/benchstat](golang.org/x/perf/cmd/benchstat).

 | name                      |     time/op |       alloc/op |     allocs/op |
|---------------------------|------------:|---------------:|--------------:|
| AgextUuidGen1Crypto-8     | 1.32µs ± 1% |     96.0B ± 0% |     3.00 ± 0% |
| AgextUuidGen1Math-8       |  273ns ± 3% |     96.0B ± 0% |     3.00 ± 0% |
| AidarkhanovNanoidV2Gen1-8 | 1.18µs ± 1% |     72.0B ± 0% |     2.00 ± 0% |
| EdgingenWuidGen1-8        | 49.1ns ± 5% |     16.0B ± 0% |     1.00 ± 0% |
| GoFlakeGen1-8             |  342ns ± 1% |     80.0B ± 0% |     5.00 ± 0% |
| GofrsUuidGen1-8           | 1.09µs ± 1% |     64.0B ± 0% |     2.00 ± 0% |
| GoogleUuidGen1Raw-8       | 1.11µs ± 1% |     64.0B ± 0% |     2.00 ± 0% |
| GoogleUuidGen1Pool-8      |  178ns ± 4% |     48.0B ± 0% |     1.00 ± 0% |
| JakehlGoidGen1-8          | 1.27µs ± 1% |      144B ± 0% |     5.00 ± 0% |
| MuyoSnoGen1-8             | 79.6ns ± 2% |     16.0B ± 0% |     1.00 ± 0% |
| OklogUlidGen1Crypto-8     | 1.17µs ± 1% |     48.0B ± 0% |     2.00 ± 0% |
| OklogUlidGen1Math-8       |  186ns ± 2% |     48.0B ± 0% |     2.00 ± 0% |
| RxXidGen1-8               |  134ns ± 1% |     24.0B ± 0% |     1.00 ± 0% |
| RwxrobUniqGen1-8          | 1.48µs ± 1% |      184B ± 0% |     7.00 ± 0% |
| SatoriGoUuidGen1-8        | 1.10µs ± 1% |     64.0B ± 0% |     2.00 ± 0% |
| TwharmonGouidLCA-8        | 1.08µs ± 1% |     16.0B ± 0% |     1.00 ± 0% |
| TwharmonGouidLCAN-8       | 1.08µs ± 2% |     16.0B ± 0% |     1.00 ± 0% |
| TwharmonGouidMCA-8        | 1.08µs ± 2% |     16.0B ± 0% |     1.00 ± 0% |
| TwharmonGouidMCAN-8       | 1.08µs ± 1% |     16.0B ± 0% |     1.00 ± 0% |
| TwharmonGouidUCA-8        | 1.07µs ± 1% |     16.0B ± 0% |     1.00 ± 0% |
| TwharmonGouidUCAN-8       | 1.08µs ± 2% |     16.0B ± 0% |     1.00 ± 0% |
