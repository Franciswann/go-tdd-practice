**Mock**

**Why?**
因為當受測對象屬於例如 time.Sleep() 這類型無法輕易監測的對象時
我們就需要inject dependency 並且implement Mock
讓它能在主程式裡跑我們要的result ('3, 2, 1 Go!')
meanwhile 在 test side 不耗時的測試 Callse