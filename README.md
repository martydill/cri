`cri` is a command-line cryptocurrency utility tool (and more importantly, an excuse to write some `golang` code).

It currently supports two commands:

`price` - fetch the price of one or more cryptocurrencies:

![Screenshot of price command](/media/price-screenshot.png?raw=true)


`chart` - generate a chart of price history for one or more cryptocurrencies:
![Screenshot of chart command](/media/chart-screenshot.png?raw=true)


The `chart` command supports a `-d` param to specify the number of days to chart:

Both commands support the `--currency` flag to change the output currency.

![Screenshot of chart command with options](/media/chart-options.png?raw=true)