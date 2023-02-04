package report

import (
	"fmt"
	"strings"
)

const PositiveResult = "在本次检索到的国内外公开发表的专利及非专利文献中，尚未发现与本项目研究内容一致的文献报道，本项目内容在国内外具备新颖性。"

const NegativeResult = "在本次检索到的国内外公开发表的专利及非专利文献中，发现与本项目研究内容一致的文献报道，本项目内容在国内外不具备新颖性。"

const ConclusionHeader = `
%d.申请人：%s
申请单位：%s
专利名称：%s
申请号：%s
申请日：%s
相似度：%s
简介：%s
`

func GenConclusionHeader(num int, userName string, depart string, patentName string,
	applyNum string, applyDate string, score string, brief string) string {
	return fmt.Sprintf(ConclusionHeader, num, userName, depart, patentName, applyNum, applyDate, score, brief)
}

const DisclaimerTemplate = `


本平台特别声明：
1、截止您提交查新报告申请之日，本查新报告充分依据提交的具体技术内容而出具。
2、本平台谨循行业勤勉尽责之精神与诚实信用之原则，以假定提供之文件资料真实有效为前提，并对之进行审慎揣度，力避所出报告记载虚假与误导陈述之内容。
3、本报告仅供参考，不作为任何法律、投资、并购、重组、买卖、许可等目的使用，未经本平台书面许可，不得另作他用。
依据主要法律（此法律系广义之说，含法律、行政法规等可作为出具报告之现行有效之立法文件）、法规、及相关规范性法律文件，通过专利搜索数据库在国内国外范围内进行检索 ，完全以文献中的事实、数据为依据，不受各种主观、客观因素影响。
`

const noveltyBase = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<style type="text/css">.mytable {
    margin: 0 auto;
    width: 620px;

}

</style>
<table class="mytable" cellspacing="0" cellpadding="0">
    <tbody>
    <tr style=";height:869px" class="firstRow">
        <td colspan="7" style="border: 1px solid windowtext; padding: 0px 7px; word-break: break-all;
        " width="553" valign="top" height="869"> <p> 报告编号：$NUMBER </p> <p> &nbsp; </p> <p>
            &nbsp; </p> <p style="text-align:center"> <strong><span style="font-size:30px;font-family:宋体">科 技 查 新 报 告</span></strong>
        </p> <p style="text-align:center"> <strong><span
                style="font-size:29px;font-family:宋体">&nbsp;</span></strong> </p> <p
                style="text-align:center"> <strong><span
                style="font-size:29px;font-family:宋体">&nbsp;</span></strong> </p> <p
                style="margin-top:16px;margin-right:0;margin-bottom:16px;margin-left:140px;text-align:left;line-height:150%">
             <strong><span style="font-size:19px;line-height:150%">项目名称：&nbsp; </span></strong><span
                style="font-size:16px;line-height:150%">$PATENT_NAME</span> </p> <p
                style="margin-top:16px;margin-right:0;margin-bottom:16px;margin-left:140px;text-align:left;line-height:150%">
             <strong><span
                style="font-size:19px;line-height:150%">委&nbsp;&nbsp;托&nbsp;&nbsp;人：&nbsp; </span></strong><span
                style="font-size:16px;line-height:150%">$USER_NAME</span> </p> <p
                style="margin-top:16px;margin-right:0;margin-bottom:16px;margin-left:140px;text-align:left;line-height:150%">
             <strong><span style="font-size:19px;line-height:150%">查新机构：&nbsp; </span></strong><span
                style="font-size:16px;line-height:150%">$Institution</span> </p> <p
                style="margin-top:16px;margin-right:0;margin-bottom:16px;margin-left:140px;text-align:left;line-height:150%">
             <strong><span style="font-size:19px;line-height:150%">完成日期：&nbsp; </span></strong><span
                style="font-size:16px;line-height:150%">$FINISH_DATE</span> </p> <p style="text-align:left">
            <strong><span style="font-size:21px">&nbsp;</span></strong> </p> <p style="text-align:left">
            <strong><span style="font-size:16px">&nbsp;</span></strong> </p> <p style="text-align:left">
            <strong><span style="font-size:16px">&nbsp;</span></strong> </p> <p style="text-align:left">
            <strong><span style="font-size:16px">&nbsp;</span></strong> </p> <p style="text-align:left">
            <strong><span style="font-size:21px">&nbsp;</span></strong> </p> <p style="text-align:left">
            <strong><span style="font-size:21px">&nbsp;</span></strong> </p> <p style="text-align:center">
            <strong><span style="font-size:16px">教育部科技发展中心</span></strong> </p> <p
                style="text-align:center"> <span style="font-size:16px">二O二三年制</span> </p>
                    </td>
    </tr>

    </tbody>

</table><p> <br/></p>
<table class="mytable" cellspacing="0" cellpadding="0">
    <tbody>
    <tr style=";height:36px" class="firstRow">
        <td rowspan="1" style="border: 1px solid windowtext; padding: 0px 7px; word-break: break-all;
        " width="77" height="36"> <p style="text-align:justify;text-justify:distribute-all-lines">
            查新项目 </p> <p style="text-align:justify;text-justify:distribute-all-lines"> 名称 </p>
                    </td>
        <td colspan="6" style="border-color: windowtext windowtext windowtext currentcolor; border-style: solid solid
            solid none; border-width: 1px 1px 1px medium; border-image: none 100%
        / 1 / 0 stretch; padding: 0px 7px; word-break: break-all;" width="476" height="36"> <p>
            中文：$PATENT_NAME </p>            </td>
    </tr>

    <tr style=";height:23px">
        <td rowspan="5" style="border-color: currentcolor windowtext windowtext; border-style: none solid solid;
            border-width: medium 1px 1px; border-image: none 100%
        / 1 / 0 stretch; padding: 0px 7px; word-break: break-all;" width="77" height="23"> <p
                style="text-align:center"> 查新机构 </p>            </td>
        <td style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none;
            border-width: medium 1px 1px medium; padding: 0px 7px;
        " width="75" height="23"> <p> 名称 </p>            </td>
        <td colspan="5" style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none
            solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;
        " width="401" height="23"> $DEPART_NAME <br/>            </td>
    </tr>

    <tr style=";height:23px">
        <td style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none;
            border-width: medium 1px 1px medium; padding: 0px 7px;
        " width="75" height="23"> <p> 通信地址 </p>            </td>
        <td colspan="3" style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none
            solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;
        " width="232" height="23"> $CONTACT_ADDR <br/>            </td>
        <td style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none;
            border-width: medium 1px 1px medium; padding: 0px 7px;
        " width="67" height="23"> <p> 邮政编码 </p>            </td>
        <td style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none;
            border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;
        " width="102" height="23"> $ZIP_CODE<br/>            </td>
    </tr>

    <tr style=";height:17px">
        <td style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none;
            border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;
        " width="75" height="17"> <p> 负责人 </p>            </td>
        <td style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none;
            border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;
        " width="93" height="17"> $MANAGER_NAME <br/>            </td>
        <td style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none;
            border-width: medium 1px 1px medium; padding: 0px 7px;
        " width="58" height="17"> <p> 电话 </p>            </td>
        <td colspan="3" style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none
            solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;
        " width="250" height="17"> $MANAGER_TEL <br/>            </td>
    </tr>

    <tr style=";height:16px">
        <td style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none;
            border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;
        " width="75" height="16"> <p> 联系人 </p>            </td>
        <td style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none;
            border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;
        " width="93" height="16"> $CONTACT_NAME <br/>            </td>
        <td style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none;
            border-width: medium 1px 1px medium; padding: 0px 7px;
        " width="58" height="16"> <p> 电话 </p>            </td>
        <td colspan="3" style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none
            solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;
        " width="250" height="16"> $CONTACT_TEL<br/>            </td>
    </tr>

    <tr style=";height:27px">
        <td style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none;
            border-width: medium 1px 1px medium; padding: 0px 7px;
        " width="75" height="27"> <p> 电子邮箱 </p>            </td>
        <td colspan="5" style="border-color: currentcolor windowtext windowtext currentcolor; border-style: none
            solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;
        " width="401" height="27"> $EMAIL <br/>            </td>
    </tr>

    <tr style=";height:107px">
        <td colspan="7" style="border-color: currentcolor windowtext windowtext; border-style: none solid solid;
            border-width: medium 1px 1px; border-image: none 100%
        / 1 / 0 stretch; padding: 0px 7px; word-break: break-all;" width="553" valign="top" height="107"> <p>
            <br> <strong><span style="font-size:20px;font-family:宋体">一、项目的科学技术要点  </span></strong></p>
        <p style="text-indent:28px">&nbsp;$TECH_POINT </p>      </td>
    </tr>

    <tr style=";height:107px">
        <td colspan="7" style="border-color: currentcolor windowtext windowtext; border-style: none solid solid;
            border-width: medium 1px 1px; border-image: none 100%
        / 1 / 0 stretch; padding: 0px 7px; word-break: break-all;" width="553" valign="top" height="107"> <p>
            <br> <strong><span style="font-size:20px;font-family:宋体">二、专利检索范围及检索策略  </span></strong>
        </p> <p style="text-indent:28px"> <strong> 检索的中文数据库 </strong></p> <p
                style="text-indent:28px"> &nbsp;$DATABASE </p> <p style="text-indent:28px"> &nbsp; </p>
        <p style="text-indent:28px"> <strong> 检索词 </strong></p> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp; $QUERY_WORD <p
                style="text-indent:28px"> &nbsp; </p> <p style="text-indent:28px"> <strong>
            检索式 </strong></p> &nbsp;&nbsp;&nbsp;&nbsp; $QUERY_EXPRESSION &nbsp; <p style="text-indent:28px"> &nbsp; </p>
                    </td>
    </tr>

    <tr style=";height:89px">
        <td colspan="7" style="border-color: currentcolor windowtext windowtext; border-style: none solid solid;
            border-width: medium 1px 1px; border-image: none 100%
        / 1 / 0 stretch; padding: 0px 7px; word-break: break-all;" width="553" valign="top" height="89"> <p>
            <br> <strong><span style="font-size:20px;font-family:宋体">  三、检索结果  </span></strong></p> <p
                style="text-indent:28px"> 依据上专利检索范围和检索式，共检索出相专利 $RELATIVE_NUM 项，其中密切相关专利 $VERY_RELATIVE_NUM
            项。 </p>&nbsp; &nbsp;&nbsp;&nbsp; $SEARCH_RESULT <p style="text-indent:28px"> &nbsp; </p>
                    </td>
    </tr>

    <tr style=";height:89px">
        <td colspan="7" style="border-color: currentcolor windowtext windowtext; border-style: none solid solid;
            border-width: medium 1px 1px; border-image: none 100%
        / 1 / 0 stretch; padding: 0px 7px; word-break: break-all;" width="553" valign="top" height="89"> <p>
            <br> <strong><span style="font-size:20px;font-family:宋体">  四、查新结论    </span></strong></p> <p
                style="text-indent:28px"> 经对检出的相关文献进行阅读、分析、对比，结论如下：</p>
        <p style="text-indent:28px"> $CONCLUSION </p>  <p> &nbsp; </p>            </td>
    </tr>

    </tbody>

</table><p> <br/></p>
</body>
</html>
`

type NoveltyTemplate struct {
	t string
}

func (t *NoveltyTemplate) Replace(old string, new string) *NoveltyTemplate {
	t.t = strings.Replace(t.String(), old, new, -1)
	return t
}

func (t *NoveltyTemplate) String() string {
	return t.t
}

func NewNoveltyTemplate() *NoveltyTemplate {
	return &NoveltyTemplate{t: noveltyBase}
}
