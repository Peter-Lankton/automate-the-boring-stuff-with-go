package main

import (
	"fmt"
	"strings"
)

func main() {
	doc := `
        In the fast-paced world of corporate business, the pursuit of optimal efficiency is paramount. 
        Companies must constantly strive to fine-tune their operational processes, fostering an environment of maximum productivity. 
        This text aims to delve into the intricacies of optimizing synergistic business operations, highlighting the various 
        mundane aspects of this essential endeavor.To begin, it is essential to emphasize the importance of aligning organizational 
        goals with strategic planning. Companies need to set clear objectives and devise comprehensive strategies to achieve them. 
        By ensuring that these objectives are communicated effectively and consistently across all departments, companies can 
        expect to see improved coordination and coherence in their operations. Once the strategic framework  a fundamental but 
        tedious undertaking that requires careful planning, continuous process improvement, meticulous monitoring, effective time 
        management, and resource allocation. While these aspects of corporate business may not be the most captivating, they are 
        integral to the success of any organization. A relentless commitment to these seemingly banal activities is what sets the 
        stage for a successful, efficient, and sustainable corporate operation.`

	if strings.Contains(doc, "aspects") {
		fmt.Println("Yes, 'aspects' is in that corporate word salad")
	}

	if !strings.Contains(doc, "orange") {
		fmt.Println("silly Gopher there is no 'orange' in that corporate word salad")
	}
	fmt.Printf("how many times is the word 'companies' used? \nans: %d", strings.Count(doc, "Companies"))

}
