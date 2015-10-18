package main

func main() {
    var myInt int
    myInt = 15

    println(myInt)

    // have to choose between 32 and 64 bit sizes
    // also have to have a decimal to force Go to realize it's a float
    var myFloat32 float32 = 15.4
    println(myFloat32)

    // short variable declaration syntax
    myString := "Hello Go"
    println(myString)

    //Complex Types
    myComplex := complex(2,3)
    println(myComplex)

    // can print the real and imaginary part seperately
    println(real(myComplex))
    println(imag(myComplex))
}
