## Go Server

# Build

To build the package, clone the GitHub repository and enter the directory, and then run the following commands:

```bash
cd src
go build
```

# Run

Run either the pre-built binary found in /build, or the executable from the commands above

```
./triumph_intern
```

# Test

To run the unit tests:

```
cd src
go test
```

# Endpoints:

GET /buy?amount={amount}&symbol={symbol}

Required parameters:
- amount (float) the amount of the first currency to purchase
- symbol (string) the currency pair, based on the list of allowed IDs from Coinbase 

Example Response

```
# curl "localhost:4000/buy?amount=1&symbol=BTC-USDT"

{
    "order": {   
        "btcAmount":1,
        "usdAmount":42230.8,
        "exchange":"kraken"
    },
    "status":"success"
}
```

----

GET /sell?amount={amount}&symbol={symbol}

Required parameters:
- amount (float) the amount of the first currency to purchase
- symbol (string) the currency pair, based on the list of allowed IDs from Coinbase 

Example Response

```
# curl "localhost:4000/sell?amount=1&symbol=BTC-USDT"

{
    "order": {   
        "btcAmount":1,
        "usdAmount":42249.1,
        "exchange":"coinbase"
    },
    "status":"success"
}
```

### Report

### Overview:

Controllers:
- buy/sell controllers handle request routing for respective endpoints

Models:
- Order model defines the actual order to be executed 
- - Ideally, the model should have sanity checks / validation for input
- - Prior the input data should initialize the Order, which will have a state and be passed to the exchange service which will process the order and updates its state

Services:
- exchangeService handles the HTTP calls to the Kraken / Coinbase API 
- - Ideally, each data source should be implemented as a separate provider that would be called instead of hard-coded in getBookData
- inputValidation handles the input validation (and ideally will further check against any other product constraints)

### **Q) If you had more time, what further improvements or new features would you add?**

#### Functional changes:

**Proper order fill**
Looking at the requested quantity and filling the order based on the available bid/asks and their respective quantities available.

**Fees**
For lower quantity transactions, fees / network can make up a non-negligible portion of the transaction, and thus affect the actual best price

#### Architectural considerations:

**Dynamic deployment port**
Change the server to read from a configuration file / env variable to determine what port to listen on

**Unit testing**
Unit tests ideally should also simulate a static response from the kraken / coinbase APIs to test that order fill / provider + price selection works properly

**Race condition**
Assuming we are actually filling the order: if two requests are made at the same time, how do we handle order fill 

#### Performance Considerations:

**Caching**
Depending on freshness constraints, use case and request rate, setting a short-lived (<1s) cache of respective API call responses would preventing hitting rate limits and improve performance

### **Q) Which parts are you most proud of? And why?**

getBestPrice() - built with consideration for how new data sources will be added (versus hard coding two sources)

### **Q) Which parts did you spend the most time with? What did you find most difficult?**

Hardest part was dealing with the data formats and how Go handles different data types

### Q) How did you find the test overall? Did you have any issues or have difficulties completing? If you have any suggestions on how we can improve the test, we'd love to hear them.

Not really.