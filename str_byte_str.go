package main

import (
		"fmt"
	)

	//var value = "satheesh"
	//var value string
	//fmt.Scanln(&value)
	func main() {
			// value := "satheesh"
				var value string
					fmt.Scanln(&value)
						fmt.Println("String:", value)

							// strings to bytes
								string_to_array := []byte(value)
									// Corresponding byte value
										fmt.Println("String To Array:", string_to_array)

											// bytes back to string
												bytes_to_string := string(string_to_array)
													// Corresponding string value
														fmt.Println("Bytes To String:", bytes_to_string)
													}
