# go语言string、int、int64互相转换以及时间转换
# string到int
    int,err:=strconv.Atoi(string)
# string到int64    
    int64, err := strconv.ParseInt(string, 10, 64)
# int到string
    string:=strconv.Itoa(int)
# int64到string    
    string:=strconv.FormatInt(int64,10) 
# 时间格式化yyyy-MM-dd hh:mm:ss
    var _startDate int64 = time.Now().Unix()  
    var startDate string = time.Unix(_startDate, 0).Format("2006-01-02 15:04:05")  
# int64转10位的时间
    t := time.Now().Unix()  
    str := strconv.FormatInt(t, 10)  
