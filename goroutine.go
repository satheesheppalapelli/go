package main

import "fmt"

func f(n int) {
		for i := 0; i < 10; i++ {
					fmt.Println(n, ":", i)
						}
					}

					func main() {
						        var value int
							fmt.Scanln(&value)
							for i := 0; i < value; i++ {
										go f(i)
											}
												var input string
													fmt.Scanln(&input)
												}
