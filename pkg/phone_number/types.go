package phonenumber

type PhoneNumber string

var (
    opCodes = []string{"982", "986", "912", "934"}
    
    formats = []string{
        `8%s\d{7}`,
        `8 (%s) \d{3}-\d{4}`,
        `8 %s \d{3} \d{2}[ ]?\d{2}`,
        `\+7%s\d{7}`,
        `\+7 (%s) \d{3}-\d{4}`,
        `\+7 %s \d{3} \d{2}[ ]?\d{2}`,
    }
)
