package util

import (
	"errors"
)

// 获取分页编号序列, 序列中非负整数表示页码, -1 表示省略, 如 [0,1,-1,8,9,10,11,12,-1,15,16] 表示 0,1,...8,9,10,11,12,...15,16
//  pageNum:    页面数量, 大于 0 的整数
//  pageIndex:  当前页码, 从 0 开始编码
func Paginator0(pageNum, pageIndex int) ([]int, error) {
	const (
		// 0,1...4,5,[6],7,8...10,11
		paginatorBeginNum = 2 // 分页开头显示的索引数目
		paginatorEndNum   = 2 // 分页结尾显示的索引数目

		pageIndexFrontNum  = 2 // 当前页面前面显示的索引数目
		pageIndexBehindNum = 2 // 当前页面后面显示的索引数目

		pageIndexRangeNum = pageIndexFrontNum + 1 + pageIndexBehindNum // 当前页面左右范围的页面个数
	)

	if pageNum < 1 {
		return nil, errors.New("pageNum < 1")
	}
	if pageIndex < 0 || pageIndex >= pageNum {
		return nil, errors.New("pageIndex out of range")
	}

	switch {
	case pageNum == 1:
		return []int{0}, nil
	case pageNum <= pageIndexRangeNum: // 不需要加省略号
		arr := make([]int, pageNum)
		for i := 0; i < pageNum; i++ {
			arr[i] = i
		}
		return arr, nil
	default: // pageNum > pageIndexRangeNum
		maxPageIndex := pageNum - 1 // maxPageIndex >= pageIndexRangeNum

		// 确定当前页面这个游标块前后的页码
		// 如 0,1...4,5,[6],7,8...10,11 里面的 4 和 8
		rangeBeginPageIndex := pageIndex - pageIndexFrontNum
		rangeEndPageIndex := pageIndex + pageIndexBehindNum
		switch {
		case rangeBeginPageIndex < 0:
			rangeBeginPageIndex = 0
			rangeEndPageIndex = pageIndexFrontNum + pageIndexBehindNum // maxPageIndex >= pageIndexRangeNum > pageIndexFrontNum + pageIndexBehindNum == rangeEndPageIndex, rangeEndPageIndex < maxPageIndex
		case rangeEndPageIndex > maxPageIndex:
			rangeEndPageIndex = maxPageIndex
			rangeBeginPageIndex = maxPageIndex - pageIndexFrontNum - pageIndexBehindNum // maxPageIndex >= pageIndexRangeNum > pageIndexFrontNum + pageIndexBehindNum, rangeBeginPageIndex > 0
		}

		if rangeBeginPageIndex <= paginatorBeginNum { // 跟前面相连
			if rangeEndPageIndex >= maxPageIndex-paginatorEndNum { // 跟后面相连
				arr := make([]int, pageNum)
				for i := 0; i < pageNum; i++ {
					arr[i] = i
				}
				return arr, nil
			} else { //跟后面不连
				arr := make([]int, 0, rangeEndPageIndex+1+1+paginatorEndNum)
				for i := 0; i <= rangeEndPageIndex; i++ {
					arr = append(arr, i)
				}
				arr = append(arr, -1)
				for i := pageNum - paginatorEndNum; i < pageNum; i++ {
					arr = append(arr, i)
				}
				return arr, nil
			}
		} else { // 跟前面不连
			if rangeEndPageIndex >= maxPageIndex-paginatorEndNum { // 跟后面相连
				arr := make([]int, 0, paginatorBeginNum+1+(pageNum-rangeBeginPageIndex))
				for i := 0; i < paginatorBeginNum; i++ {
					arr = append(arr, i)
				}
				arr = append(arr, -1)
				for i := rangeBeginPageIndex; i < pageNum; i++ {
					arr = append(arr, i)
				}
				return arr, nil
			} else { //跟后面不连
				arr := make([]int, 0, paginatorBeginNum+1+pageIndexRangeNum+1+paginatorEndNum)
				for i := 0; i < paginatorBeginNum; i++ {
					arr = append(arr, i)
				}
				arr = append(arr, -1)
				for i := rangeBeginPageIndex; i <= rangeEndPageIndex; i++ {
					arr = append(arr, i)
				}
				arr = append(arr, -1)
				for i := pageNum - paginatorEndNum; i < pageNum; i++ {
					arr = append(arr, i)
				}
				return arr, nil
			}
		}
	}
}

// 获取分页编号序列, 序列中正整数表示页码, -1 表示省略, 如 [1,2,-1,8,9,10,11,12,-1,15,16] 表示 1,2,...8,9,10,11,12,...15,16
//  pageNum:    页面数量, 大于 0 的整数
//  pageIndex:  当前页码, 从 1 开始编码
func Paginator1(pageNum, pageIndex int) ([]int, error) {
	pageIndex--
	arr, err := Paginator0(pageNum, pageIndex)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != -1 {
			arr[i]++
		}
	}
	return arr, nil
}

// 获取分页编号序列, 序列中非负整数表示页码, -1 表示省略, 如 [0,1,-1,8,9,10,11,12,-1,15,16] 表示 0,1,...8,9,10,11,12,...15,16
//  totalNum:   总的记录数量, 不是页面数量
//  numPerPage: 每页显示的数量
//  pageIndex:  当前页码, 从 0 开始编码
func Paginator0Ex(totalNum, numPerPage, pageIndex int) (arr []int, pageNum int, err error) {
	if totalNum < 0 {
		err = errors.New("totalNum < 0")
		return
	}
	if numPerPage <= 0 {
		err = errors.New("numPerPage <= 0")
		return
	}

	if totalNum == 0 {
		pageNum = 1
	} else { // totalNum > 0
		pageNum = totalNum / numPerPage
		if pageNum*numPerPage < totalNum {
			pageNum++
		}
	}

	if arr, err = Paginator0(pageNum, pageIndex); err != nil {
		return nil, 0, err
	}
	return
}

// 获取分页编号序列, 序列中正整数表示页码, -1 表示省略, 如 [1,2,-1,8,9,10,11,12,-1,15,16] 表示 1,2,...8,9,10,11,12,...15,16
//  totalNum:   总的记录数量, 不是页面数量
//  numPerPage: 每页显示的数量
//  pageIndex:  当前页码, 从 1 开始编码
func Paginator1Ex(totalNum, numPerPage, pageIndex int) (arr []int, pageNum int, err error) {
	if totalNum < 0 {
		err = errors.New("totalNum < 0")
		return
	}
	if numPerPage <= 0 {
		err = errors.New("numPerPage <= 0")
		return
	}

	if totalNum == 0 {
		pageNum = 1
	} else { // totalNum > 0
		pageNum = totalNum / numPerPage
		if pageNum*numPerPage < totalNum {
			pageNum++
		}
	}

	if arr, err = Paginator1(pageNum, pageIndex); err != nil {
		return nil, 0, err
	}
	return
}
