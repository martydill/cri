`cri` is a command-line cryptocurrency utility tool (and more importantly, an excuse to write some `golang` code).

It currently supports two commands:

`price` - fetch the price of one or more cryptocurrencies:
```
MacBook:cri martydill$ go run cri price bitcoin,litecoin,ethereum,tether
bitcoin: 41263.000000
litecoin: 71.780000
ethereum: 2189.820000
tether: 0.997957
```

`chart` - generate a chart of price history for a given cryptocurrency
```
 go run cri chart bitcoin
 44035 ┤                                                                            ╭─╮
 42247 ┤                                                                         ╭──╯ │
 40459 ┤                                                                        ╭╯    ╰
 38671 ┤                                                                      ╭─╯
 36883 ┤                                                  ╭───╮   ╭───────────╯
 35095 ┤                                            ╭╮ ╭╮╭╯   ╰───╯
 33307 ┤                                    ╭───────╯╰─╯╰╯
 31519 ┤                                    │
 29731 ┤                                 ╭──╯
 27943 ┤                    ╭──╮     ╭───╯
 26155 ┼────────────────────╯  ╰─────╯
 ```

The `chart` command supports a `-d` param to specify the number of days to chart:

Both commands support the `--currency` flag to change the output currency.

```
go run cri chart ethereum -d 22 --currency jpy
 341015 ┤                                                                ╭──────────╮
 336354 ┤                                                             ╭──╯          ╰╮
 331694 ┤                                                        ╭────╯              ╰─╮
 327033 ┤                                                    ╭───╯                     ╰
 322373 ┤                                                   ╭╯
 317713 ┤                                                 ╭─╯
 313052 ┤                ╭─────╮                        ╭─╯
 308392 ┤            ╭───╯     ╰──╮               ╭─────╯
 303731 ┤         ╭──╯            ╰────────╮   ╭──╯
 299071 ┤ ╭───────╯                        ╰───╯
 294410 ┼─╯
 ```