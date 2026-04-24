// <<<<<<<<< VERSION ONE >>>>>>>>

func isValid(s string) bool {
    if len(s) == 1 {
        return false
    }

    // open := []string{"(", "[", "{"}
    // close := []string{")", "]", "}"}
    for i := 0; i < len(s); i++ {
        for j := 0; j < len(s)-1; j++ {
            if string(s[i]) == "(" && string(s[j]) == ")"{
                if string(s[i+1]) != ")" || string(s[i+2]) == ")" {
                }
                return true
            }
            if string(s[i]) == "{" && string(s[j]) == "}" {
                if string(s[i+1]) != "}" || string(s[i+2]) == "}" {
                }
                return true
            }
            if string(s[i]) == "[" && string(s[j]) == "]" {
                if string(s[i+1]) != "[" && string(s[i+2]) != "[" {
                    return false
                } else {
                    return true
                }
            }
        }
        
  }
  return false
}

//<<<<<<< VERSION TWO >>>>>>>>>