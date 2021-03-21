# Football-data-go-parallelism

Rob pike one of the creators of golang programming said, "concurrency != parallelism". "concurrency is dealing with a lot of things at once, parallel is doing a lot of things at once, one is about structure, the other is about execution"

This implementation aims to fetch a batched football data from an API parallelly over pages of data instead of fetching it per page
