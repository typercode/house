package main

import (
	"flag"
	"fmt"
)

/**
计算等额本金、提前还款（减少月还款额）

数据校对：

https://www.cmbchina.com/CmbWebPubInfo/Cal_Loan_Per.aspx?chnl=dkjsq
*/

func main() {

	_totalAmount := flag.Float64("totalAmount", 1200000, "设置总贷款额")

	_count := flag.Float64("count", 12*30, "设置总还款期数（月）")

	_monthlyInterestRate := flag.Float64("monthlyInterestRate", 0.00370833, "设置月利率")

	_earlyRepaymentPeriod := flag.Float64("earlyRepaymentPeriod", 0, "设置第几期提前还款")

	_advancePaymentAmount := flag.Float64("advancePaymentAmount", 0, "设置提前还款金额")

	flag.Parse()

	var totalAmount = *_totalAmount

	var count = *_count

	//月利率
	var monthlyInterestRate = *_monthlyInterestRate

	var earlyRepaymentPeriod = *_earlyRepaymentPeriod

	var advancePaymentAmount = *_advancePaymentAmount

	fmt.Printf("贷款额 %v\n", totalAmount)
	fmt.Printf("还款期数 %v\n", count)
	fmt.Printf("年利率 %v\n", monthlyInterestRate*12)
	fmt.Printf("第几期提前还款 %v\n", earlyRepaymentPeriod)
	fmt.Printf("设置提前还款金额 %v\n", advancePaymentAmount)

	//总利息
	var totalInterest float64 = 0

	//已还本金
	var repaidAmount float64 = 0

	var i float64 = 1

	for ; i <= count; i++ {
		if totalAmount <= repaidAmount {
			fmt.Printf("月供供完啦！！！")
			break
		}
		//当期需还利息
		var currentInterest = (totalAmount - repaidAmount) * monthlyInterestRate

		//当期需还款本金

		var currentPrincipal = (totalAmount - repaidAmount) / (count + 1 - i)

		totalInterest = totalInterest + currentInterest

		var currentPrincipalAndInterest = currentPrincipal + currentInterest

		fmt.Printf("第%v期，月供: %v，月供本金：%v,月供利息：%v\n", i, currentPrincipalAndInterest, currentPrincipal, currentInterest)

		if i == earlyRepaymentPeriod { //第i期的时候多还钱了
			repaidAmount = repaidAmount + advancePaymentAmount
		}
		repaidAmount = repaidAmount + currentPrincipal
		if repaidAmount > totalAmount {
			fmt.Printf("多还钱啦！！！%v\n", repaidAmount-totalAmount)
			repaidAmount = totalAmount
		}
	}

	fmt.Printf("总利息：%v\n", totalInterest)
	fmt.Printf("已还本金：%v\n", repaidAmount)
}
