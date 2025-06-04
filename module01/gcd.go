package module01

// Another way to solve GCD is using the Euclidean algorithm. We won't dive into the math of it, but what makes this algorithm really nice is both its simplicity and the fact that we don't actually need to factor numbers using primes.

// Here is how it works:

// Given two numbers, A and B:

// Step 1: If B == 0, return A
// Step 2: A becomes B, and B becomes the remainder of dividing A by B
// `a, b = b, a % b`
// Step 3: Go to step 1

// *Note: It doesn't matter if `A > B` or `A < B`.*

// This is a really good algorithm to know because GCD is often used in algorithmic contests. For instance, GCD was used in a qualifier for the most recent (2019) Google Code Jam. It is also a relatively simple algorithm that we can look at and then try to translate into code while exploring both recursive and iterative solutions (which we will do in a later module on recursion).

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	for b > 0 {
		a, b = b, a%b
	}
	return a

}
