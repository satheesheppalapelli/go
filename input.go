package main

import (
		"fmt"
			"reflect"
				"strconv"
			)

			func whichType(x interface{}) {
					// var x interface{}
						switch v := x.(type) {
								case int:
											fmt.Println("int:", v)
												case float64:
															fmt.Println("float64:", v)
																case string:
																			fmt.Println("string:", v)
																				case bool:
																							fmt.Println("bool:", v)
																								default:
																											fmt.Println("unknown")
																												}
																											}

																											func main() {
																													fmt.Println("Enter Value: ")
																														// Read Value
																															var s string
																																fmt.Scanf("%s", &s)

																																	st := reflect.TypeOf(s).Kind()

																																		// covert string to int
																																			// using strconv.ParseInt
																																				// i, err := strconv.ParseInt(s, 10, 64)
																																					// if err == nil {
																																						//	fmt.Println(i)
																																							// }

																																								// using strconv.Atoi
																																									if i, err := strconv.Atoi(s); err == nil {
																																												// convert string to int
																																														whichType(i)
																																															} else if b, err := strconv.ParseBool(s); err == nil {
																																																		// convert string to bool
																																																				whichType(b)
																																																					} else if f, err := strconv.ParseFloat(s, 64); err == nil {
																																																								// convert string to float
																																																										whichType(f)
																																																											} else if st == reflect.String {
																																																														// to check value is string
																																																																whichType(s)
																																																																	} else {
																																																																				// unkown value
																																																																						fmt.Printf("Enter Valid Value")
																																																																							}

																																																																								//reader := bufio.NewReader(os.Stdin)
																																																																									//fmt.Print("Enter text: ")
																																																																										//text, _ := reader.ReadString('\n')
																																																																											//whichType(text)

																																																																												//in := reflect.TypeOf(text).Kind()
																																																																													//fl := reflect.TypeOf(text).Kind()
																																																																														//bl := reflect.TypeOf(text).Kind()
																																																																															//st := reflect.TypeOf(text).Kind()

																																																																																//if in == reflect.Int {
																																																																																	//	whichType(text)
																																																																																		//} else if fl == reflect.Float64 {
																																																																																			//	whichType(text)
																																																																																				//} else if bl == reflect.Bool {
																																																																																					//	whichType(text)
																																																																																						//} else if st == reflect.String {
																																																																																							//	whichType(text)
																																																																																								//} else {
																																																																																									//	fmt.Printf("Enter Valid Value")
																																																																																										//}

																																																																																									}

