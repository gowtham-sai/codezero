# Code Zero

This project is to save time for developers and testers inside GOJEK by mocking external dependecies
on wire and returning the response they want. This will testers and developers without even changing the integration 
servers down or changing the timeout for them.
 
How do we produce the response of Traps (Transport Estimation Service)
when GOPAY is down without actually making GOPAY server down in integration?


## How this works?
Current plan is that, you can define all the external scenarios, and you can define a test file consisting
of all the scenarios you named of external dependencies.  


#### Exmaple
You define a scenario named gopay_disabled and mock the gopay service response, and in your test file scenarios,
you can say Scenario name goapy_disabled and you can use this along with other scenarios combined. With these, you 
can genrate scenario of gopay disabled and vouchers enabled without actually making any changes to integration servers
from pipeline.  


