package algo


func getMax(operations []string) []int32 { 

    // Write your code here 

    stack := make([]int32, 0) 

    result := make([]int32, 0) 

     

    for _, v := range operations { 

        innerVal := strings.Split(v, " ") 

        if innerVal[0] == "1" { 

            intVal, _ := strconv.Atoi(innerVal[1]) 

            stack = append(stack, int32(intVal)) 

        } else if innerVal[0] == "2" { 

            // remove the top element 

            stack = stack[:len(stack)-1] 

        } else if innerVal[0] == "3" { 

            // find the maximum 

            var temp int32 

            for i:=0; i<len(stack); i++ { 

                if stack[i] > temp { 

                    temp = stack[i] 

                } 

            } 

            // append the max to the result  

            result = append(result, int32(temp)) 

        } 

         

    } 

    return result 

} 



  

func Autocorrect(text string) string { 

  re := regexp.MustCompile(`(?i)\b(u|you+)\b`) 

  return re.ReplaceAllString(text, "your client") 

} 

 
 

func PackArray(integers []int) int { 

	// store the number of iterations 
  
	count := 1 
  
	 
  
	// loop through the integers until it contains one element 
  
	for len(integers) > 1 { 
  
	  result := []int{} 
  
	  for i := 0; i < len(integers)-1; i += 2 { 
  
		if count % 2 == 1 { 
  
		  result = append(result, integers[i]+integers[i+1]) 
  
		} else { 
  
		  result = append(result, integers[i]*integers[i+1]) 
  
		} 
  
	  } 
  
	  integers = result 
  
	  count++ 
  
	} 
  
	return integers[0] 
  
  } 

  func PackArray(integers []int) int { 

	return helper(integers, 1)[0] 
  
  } 
  
	
  
  func helper(arr []int, count int) []int { 
  
	if len(arr) == 1 { 
  
	  return arr 
  
	} 
  
	result := []int{} 
  
	for i := 0; i < len(arr)-1; i += 2 { 
  
	  if count % 2 == 1 { 
  
		result = append(result, arr[i]+arr[i+1]) 
  
	  } else { 
  
		 result = append(result, arr[i]*arr[i+1]) 
  
	  } 
  
	} 
  
	return helper(result, count+1) 
  
  } 

  func getMax(operations []string) []int32 { 

    // Write your code here 

    stack := make([]int32, 0) 

    result := make([]int32, 0) 

     

    for _, v := range operations { 

        innerVal := strings.Split(v, " ") 

        if innerVal[0] == "1" { 

            intVal, _ := strconv.Atoi(innerVal[1]) 

            stack = append(stack, int32(intVal)) 

        } else if innerVal[0] == "2" { 

            // remove the top element 

            stack = stack[:len(stack)-1] 

        } else if innerVal[0] == "3" { 

            // find the maximum 

            var temp int32 

            for i:=0; i<len(stack); i++ { 

                if stack[i] > temp { 

                    temp = stack[i] 

                } 

            } 

            // append the max to the result  

            result = append(result, int32(temp)) 

        } 

         

    } 

    return result 

} 


func countingValleys(steps int32, path string) int32 { 

    // Write your code here 

    var level, valley int32 

    for _, p := range path { 

        if p == 'U' { 

            level++ 

        } 

        if p == 'D' { 

            level-- 

        } 

        if level == 0 && p == 'U' { 

            valley++ 

        } 

    } 

    return valley 

} 





 

func rotateLeft(d int32, arr []int32) []int32 { 

    // Write your code here 

    result := arr[d:] 

    arr = arr[:d] 

    result = append(result, arr...) 

    return result 

} 

 



func rotateLeft(d int32, arr []int32) []int32 { 

    // Write your code here 

    result := make([]int32, len(arr)) 

    total := int32(len(arr)) 

    for idx, val := range arr { 

        innerV := total + (int32(idx)-d) 

        result[innerV%total] = val 

    } 

    return result 

 } 

 //Enter your code here. Read input from STDIN. Print output to STDOUT 

result := make([]int,0) 

reader := bufio.NewReader(os.Stdin) 

inputReader, _, err := reader.ReadLine() 

if err != nil { 

     log.Fatal(err) 

} 

n, err := strconv.Atoi(strings.TrimSpace(strings.Trim(string(inputReader), "\r\n"))) 

if err != nil { 

     log.Printf("An error occurred: %v", err.Error()) 

return 

} 

for i := 1; i <= n; i++ { 

    inputReader, _, err := reader.ReadLine()    if err != nil { 

        log.Fatal(err) 

    } 

    if inputReader[0] == '1' { 

        val, _ := strconv.Atoi(strings.TrimSpace(strings.Trim(string(inputReader[1:]), "\r\n"))) 

        result = append(result, val) 

    } else if inputReader[0] == '2' { 

        result = result[1:] 

    } else if inputReader[0] == '3' { 

        fmt.Printf("%d\n", result[0]) 

    } 

} 

} 
 