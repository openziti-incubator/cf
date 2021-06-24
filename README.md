# Cf

Cf is a `golang` library that offers an opinionated, idiomatic approach for building composable, extensible software frameworks.

Cf is the result of years of work building next-generation network overlay software (see [Ziti](https://ziti.dev/)). Ziti is composed of layered, orthogonal frameworks that encapsulate aspects of the overlay architecure. Ziti is designed to be extensible, and can be configured to work in numerous ways depending on the needs of a specific deployment. Many of the components may or may not exist in a Ziti configuration, and each of the components likely requires a high degree of configurability. Cf helps to manage this kind of complexity.

Cf is spiritually similar to classic "inversion of control" containers as used in languages like Java, but done in an idiomatic `golang` style.
