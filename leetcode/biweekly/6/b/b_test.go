// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[1,0,1,0,1]`, 
			`1`,
		},
		{
			`[0,0,0,1,0]`, 
			`0`,
		},
		{
			`[1,0,1,0,1,0,0,1,1,0,1]`, 
			`3`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minSwaps, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-6/problems/minimum-swaps-to-group-all-1s-together/
