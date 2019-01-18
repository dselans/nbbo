# Exercise w/ NBBO

* Start a TCP server that emits "ticker" messages for symbols on various exchanges with bids and offers
* Open a connection to the local server; listen for messages that look like quotes
* Parse and compare incoming quote against local quotes

```
$ go run main.go
Listening on localhost:7777
>> Server generated quote: Q|ibm|nyse|47|59
DING DING! Symbol: ibm || Best bid: 47 on 'nyse' || Best offer: 59 on 'nyse'
>> Server generated quote: Q|ibm|nyse|25|40
DING DING! Symbol: ibm || Best bid: 25 on 'nyse' || Best offer: 59 on 'nyse'
>> Server generated quote: Q|dell|nyse|94|11
DING DING! Symbol: dell || Best bid: 94 on 'nyse' || Best offer: 11 on 'nyse'
>> Server generated quote: Q|google|bats|28|74
DING DING! Symbol: google || Best bid: 28 on 'bats' || Best offer: 74 on 'bats'
>> Server generated quote: Q|walmart|bats|37|6
DING DING! Symbol: walmart || Best bid: 37 on 'bats' || Best offer: 6 on 'bats'
>> Server generated quote: Q|walmart|nyse|28|58
DING DING! Symbol: walmart || Best bid: 28 on 'nyse' || Best offer: 58 on 'nyse'
>> Server generated quote: Q|walmart|nasdaq|87|88
DING DING! Symbol: walmart || Best bid: 28 on 'nyse' || Best offer: 88 on 'nasdaq'
>> Server generated quote: Q|google|nyse|41|8
>> Server generated quote: Q|walmart|bats|29|56
>> Server generated quote: Q|ibm|nasdaq|85|26
>> Server generated quote: Q|ibm|nasdaq|94|63
DING DING! Symbol: ibm || Best bid: 25 on 'nyse' || Best offer: 63 on 'nasdaq'
>> Server generated quote: Q|ibm|nyse|78|24
>> Server generated quote: Q|walmart|nyse|57|21
>> Server generated quote: Q|ibm|nyse|0|5
DING DING! Symbol: ibm || Best bid: 0 on 'nyse' || Best offer: 63 on 'nasdaq'
>> Server generated quote: Q|dell|nyse|3|55
DING DING! Symbol: dell || Best bid: 3 on 'nyse' || Best offer: 55 on 'nyse'
>> Server generated quote: Q|walmart|nyse|5|56
DING DING! Symbol: walmart || Best bid: 5 on 'nyse' || Best offer: 88 on 'nasdaq'
>> Server generated quote: Q|google|bats|61|2
>> Server generated quote: Q|walmart|bats|63|76
>> Server generated quote: Q|google|bats|47|94
DING DING! Symbol: google || Best bid: 28 on 'bats' || Best offer: 94 on 'bats'
>> Server generated quote: Q|ibm|nasdaq|96|20
>> Server generated quote: Q|walmart|nyse|37|33
>> Server generated quote: Q|ibm|nasdaq|33|43
^Csignal: interrupt
``` 
