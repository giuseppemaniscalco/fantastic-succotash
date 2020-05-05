# Spark Vending machine

Our company is growing very fast with all the new acquisitions. We are trying to cope with all the demand our employees have for food and beverages.

We estimated that it would be too hard to acquire multiple vending machines and suite them to our needs, therefore we decided to build our own, smarter ones.

Together we must design such a machine based on the following requirements:

- It will accept as payment coins of 1,5,10,25 Sparks.

- We have an inventory of available coins at any point in time, which later can help us see if change can be returned.

- We also have an inventory of all available products, which can help us see which products are still available for purchase at any point in time.

- A employee can choose from different products SparklingWater(25 Sparks), SparkPasta(35 Sparks), Sparksoda(45 Sparks).

- Before returning the selected products, we will validate if enough currency is inserted to accommodate the price of the selected products.

- Before returning the selected products, we also need to check if we can supply the employee with the necessary change from our currency inventory (if needed)

- Upon confirmation of the purchase the vending machine will return the selected products.

- The product will be removed from the product inventory.

- Additionally to the product the machine will offer any remaining change back to the employee if the amount is larger than needed.

- If an employee will change his mind about the selected products, he can cancel and just get a refund of all his Sparks.

- Since products can be depleted, we can have a simple resupply functionality.

- (**BONUS POINTS**) A vending machine should be able to serve requests for more than one user simultaneously. Our hardware will have multiple coin input slots and selection, our software needs to cope with that. We need to avoid having two of our colleagues trying to acquire the last of our SparklingWater at the same time. We need to ensure at least one of them will get it, right ?

### Development

- Use the language you feel most comfortable with

- You can leverage any means on information: google it, old projects, copy-paste code e.t.c

- We need to ensure that any change we do is working, our employees would be disappointed if we roll out a non-functioning unit, and we do value their opinion. We should employ TDD for this.

- Rome was not built in one day, and we do not expect us to be able to solve all the necessary requirements. But anything we deliver has to be in good shape and quality hence TDD will help.

### Gotchas:

- Maybe the machine does not have enough funds to give back change, or the remaining coins can't accommodate the needed change amount. How do we solve this ?

- How do we ensure proper and efficient locking of our products for multiple users ?