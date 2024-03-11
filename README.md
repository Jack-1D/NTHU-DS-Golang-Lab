# NTHU-DS-Golang-Lab
Lab1 for distributed system

筆記：
select和switch是不一樣的。switch會sequential的去檢查每一個條件，select是blocking的，並且當多個條件同時滿足時，會隨機挑一個來做

因此不能這樣寫：

![](https://hackmd.io/_uploads/By_NW5n6a.png)
