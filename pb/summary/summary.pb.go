/*
Copyright The TestGrid Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: summary.proto

package summary

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TestInfo_Trend int32

const (
	TestInfo_UNKNOWN   TestInfo_Trend = 0
	TestInfo_NO_CHANGE TestInfo_Trend = 1
	TestInfo_UP        TestInfo_Trend = 2
	TestInfo_DOWN      TestInfo_Trend = 3
)

var TestInfo_Trend_name = map[int32]string{
	0: "UNKNOWN",
	1: "NO_CHANGE",
	2: "UP",
	3: "DOWN",
}

var TestInfo_Trend_value = map[string]int32{
	"UNKNOWN":   0,
	"NO_CHANGE": 1,
	"UP":        2,
	"DOWN":      3,
}

func (x TestInfo_Trend) String() string {
	return proto.EnumName(TestInfo_Trend_name, int32(x))
}

func (TestInfo_Trend) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{1, 0}
}

type DashboardTabSummary_TabStatus int32

const (
	DashboardTabSummary_NOT_SET DashboardTabSummary_TabStatus = 0
	DashboardTabSummary_UNKNOWN DashboardTabSummary_TabStatus = 1
	DashboardTabSummary_PASS    DashboardTabSummary_TabStatus = 2
	DashboardTabSummary_FAIL    DashboardTabSummary_TabStatus = 3
	DashboardTabSummary_FLAKY   DashboardTabSummary_TabStatus = 4
	DashboardTabSummary_STALE   DashboardTabSummary_TabStatus = 5
	DashboardTabSummary_BROKEN  DashboardTabSummary_TabStatus = 6
)

var DashboardTabSummary_TabStatus_name = map[int32]string{
	0: "NOT_SET",
	1: "UNKNOWN",
	2: "PASS",
	3: "FAIL",
	4: "FLAKY",
	5: "STALE",
	6: "BROKEN",
}

var DashboardTabSummary_TabStatus_value = map[string]int32{
	"NOT_SET": 0,
	"UNKNOWN": 1,
	"PASS":    2,
	"FAIL":    3,
	"FLAKY":   4,
	"STALE":   5,
	"BROKEN":  6,
}

func (x DashboardTabSummary_TabStatus) String() string {
	return proto.EnumName(DashboardTabSummary_TabStatus_name, int32(x))
}

func (DashboardTabSummary_TabStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{4, 0}
}

// Summary of a failing test.
type FailingTestSummary struct {
	// Display name of the test.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Name of the test. E.g., the target for tests in Sponge.
	TestName string `protobuf:"bytes,2,opt,name=test_name,json=testName,proto3" json:"test_name,omitempty"`
	// First build ID at which the test failed.
	FailBuildId string `protobuf:"bytes,3,opt,name=fail_build_id,json=failBuildId,proto3" json:"fail_build_id,omitempty"`
	// Timestamp for the first cycle in which the test failed.
	FailTimestamp float64 `protobuf:"fixed64,4,opt,name=fail_timestamp,json=failTimestamp,proto3" json:"fail_timestamp,omitempty"`
	// Last build ID at which the test passed.
	PassBuildId string `protobuf:"bytes,5,opt,name=pass_build_id,json=passBuildId,proto3" json:"pass_build_id,omitempty"`
	// Timestamp for the last cycle in which the test passed.
	PassTimestamp float64 `protobuf:"fixed64,6,opt,name=pass_timestamp,json=passTimestamp,proto3" json:"pass_timestamp,omitempty"`
	// Number of times the test has failed.
	FailCount int32 `protobuf:"varint,7,opt,name=fail_count,json=failCount,proto3" json:"fail_count,omitempty"`
	// Link to search for build changes.
	BuildLink string `protobuf:"bytes,8,opt,name=build_link,json=buildLink,proto3" json:"build_link,omitempty"`
	// Text for option to search for build changes.
	BuildLinkText string `protobuf:"bytes,9,opt,name=build_link_text,json=buildLinkText,proto3" json:"build_link_text,omitempty"`
	// Text to display for link to search for build changes.
	BuildUrlText string `protobuf:"bytes,10,opt,name=build_url_text,json=buildUrlText,proto3" json:"build_url_text,omitempty"`
	// Text for failure statuses associated with this test.
	FailureMessage string `protobuf:"bytes,11,opt,name=failure_message,json=failureMessage,proto3" json:"failure_message,omitempty"`
	// List of bug IDs for bugs associated with this test.
	LinkedBugs []string `protobuf:"bytes,12,rep,name=linked_bugs,json=linkedBugs,proto3" json:"linked_bugs,omitempty"`
	// A link to the first build in which the test failed.
	FailTestLink string `protobuf:"bytes,13,opt,name=fail_test_link,json=failTestLink,proto3" json:"fail_test_link,omitempty"`
	// A link to the latest build in which the test failed.
	LatestFailTestLink string `protobuf:"bytes,17,opt,name=latest_fail_test_link,json=latestFailTestLink,proto3" json:"latest_fail_test_link,omitempty"`
	// The test ID for the latest test failure. (Does not indicate the failure is
	// 'over', just the latest test failure we found.)
	LatestFailBuildId string `protobuf:"bytes,14,opt,name=latest_fail_build_id,json=latestFailBuildId,proto3" json:"latest_fail_build_id,omitempty"`
	// Maps (property name):(property value) for arbitrary alert properties.
	Properties map[string]string `protobuf:"bytes,15,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of IDs for issue hotlists related to this failure.
	HotlistIds []string `protobuf:"bytes,16,rep,name=hotlist_ids,json=hotlistIds,proto3" json:"hotlist_ids,omitempty"`
	// Dynamic email list, route email alerts to these instead of the configured defaults.
	EmailAddresses       []string `protobuf:"bytes,18,rep,name=email_addresses,json=emailAddresses,proto3" json:"email_addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FailingTestSummary) Reset()         { *m = FailingTestSummary{} }
func (m *FailingTestSummary) String() string { return proto.CompactTextString(m) }
func (*FailingTestSummary) ProtoMessage()    {}
func (*FailingTestSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{0}
}

func (m *FailingTestSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailingTestSummary.Unmarshal(m, b)
}
func (m *FailingTestSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailingTestSummary.Marshal(b, m, deterministic)
}
func (m *FailingTestSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailingTestSummary.Merge(m, src)
}
func (m *FailingTestSummary) XXX_Size() int {
	return xxx_messageInfo_FailingTestSummary.Size(m)
}
func (m *FailingTestSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_FailingTestSummary.DiscardUnknown(m)
}

var xxx_messageInfo_FailingTestSummary proto.InternalMessageInfo

func (m *FailingTestSummary) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *FailingTestSummary) GetTestName() string {
	if m != nil {
		return m.TestName
	}
	return ""
}

func (m *FailingTestSummary) GetFailBuildId() string {
	if m != nil {
		return m.FailBuildId
	}
	return ""
}

func (m *FailingTestSummary) GetFailTimestamp() float64 {
	if m != nil {
		return m.FailTimestamp
	}
	return 0
}

func (m *FailingTestSummary) GetPassBuildId() string {
	if m != nil {
		return m.PassBuildId
	}
	return ""
}

func (m *FailingTestSummary) GetPassTimestamp() float64 {
	if m != nil {
		return m.PassTimestamp
	}
	return 0
}

func (m *FailingTestSummary) GetFailCount() int32 {
	if m != nil {
		return m.FailCount
	}
	return 0
}

func (m *FailingTestSummary) GetBuildLink() string {
	if m != nil {
		return m.BuildLink
	}
	return ""
}

func (m *FailingTestSummary) GetBuildLinkText() string {
	if m != nil {
		return m.BuildLinkText
	}
	return ""
}

func (m *FailingTestSummary) GetBuildUrlText() string {
	if m != nil {
		return m.BuildUrlText
	}
	return ""
}

func (m *FailingTestSummary) GetFailureMessage() string {
	if m != nil {
		return m.FailureMessage
	}
	return ""
}

func (m *FailingTestSummary) GetLinkedBugs() []string {
	if m != nil {
		return m.LinkedBugs
	}
	return nil
}

func (m *FailingTestSummary) GetFailTestLink() string {
	if m != nil {
		return m.FailTestLink
	}
	return ""
}

func (m *FailingTestSummary) GetLatestFailTestLink() string {
	if m != nil {
		return m.LatestFailTestLink
	}
	return ""
}

func (m *FailingTestSummary) GetLatestFailBuildId() string {
	if m != nil {
		return m.LatestFailBuildId
	}
	return ""
}

func (m *FailingTestSummary) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *FailingTestSummary) GetHotlistIds() []string {
	if m != nil {
		return m.HotlistIds
	}
	return nil
}

func (m *FailingTestSummary) GetEmailAddresses() []string {
	if m != nil {
		return m.EmailAddresses
	}
	return nil
}

// Metrics about a specific test, i.e. passes, fails, total runs, etc.
// Next ID: 12
type TestInfo struct {
	// The display name of the test, typically what is shown for each row in TestGrid
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The total number of test runs not including runs failed due to infrastructure
	// failures.
	TotalNonInfraRuns int32 `protobuf:"varint,2,opt,name=total_non_infra_runs,json=totalNonInfraRuns,proto3" json:"total_non_infra_runs,omitempty"`
	// The number of passed test runs not including runs failed due to
	// infrastructure failures.
	PassedNonInfraRuns int32 `protobuf:"varint,3,opt,name=passed_non_infra_runs,json=passedNonInfraRuns,proto3" json:"passed_non_infra_runs,omitempty"`
	// The number of failed test runs not including runs failed due to
	// infrastructure failures.
	FailedNonInfraRuns int32 `protobuf:"varint,4,opt,name=failed_non_infra_runs,json=failedNonInfraRuns,proto3" json:"failed_non_infra_runs,omitempty"`
	// The number of failed test runs specifically due to infrastructure
	// failures.
	FailedInfraRuns int32 `protobuf:"varint,5,opt,name=failed_infra_runs,json=failedInfraRuns,proto3" json:"failed_infra_runs,omitempty"`
	// The total number of all runs, including failures due to infrastructure
	TotalRunsWithInfra int32 `protobuf:"varint,6,opt,name=total_runs_with_infra,json=totalRunsWithInfra,proto3" json:"total_runs_with_infra,omitempty"`
	// Any other type of runs not included above.
	OtherRuns int32 `protobuf:"varint,7,opt,name=other_runs,json=otherRuns,proto3" json:"other_runs,omitempty"`
	// The flakiness of the test, measured out of 100
	Flakiness float32 `protobuf:"fixed32,8,opt,name=flakiness,proto3" json:"flakiness,omitempty"`
	// The flakiness of the test from previous intervals
	PreviousFlakiness []float32 `protobuf:"fixed32,10,rep,packed,name=previous_flakiness,json=previousFlakiness,proto3" json:"previous_flakiness,omitempty"`
	// The change of flakiness based on the last interval's flakiness
	// e.g. if last interval the flakiness was 50, and now it's 75, the
	// trend is UP. A trend of NO_CHANGE means last week and this week were
	// exactly the same. The interval is set by each tab's config, with
	// a default of 7 days.
	ChangeFromLastInterval TestInfo_Trend `protobuf:"varint,9,opt,name=change_from_last_interval,json=changeFromLastInterval,proto3,enum=TestInfo_Trend" json:"change_from_last_interval,omitempty"`
	// A map of infra failure name to the count of that failure for the interval.
	InfraFailures        map[string]int32 `protobuf:"bytes,11,rep,name=infra_failures,json=infraFailures,proto3" json:"infra_failures,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TestInfo) Reset()         { *m = TestInfo{} }
func (m *TestInfo) String() string { return proto.CompactTextString(m) }
func (*TestInfo) ProtoMessage()    {}
func (*TestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{1}
}

func (m *TestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestInfo.Unmarshal(m, b)
}
func (m *TestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestInfo.Marshal(b, m, deterministic)
}
func (m *TestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestInfo.Merge(m, src)
}
func (m *TestInfo) XXX_Size() int {
	return xxx_messageInfo_TestInfo.Size(m)
}
func (m *TestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TestInfo proto.InternalMessageInfo

func (m *TestInfo) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TestInfo) GetTotalNonInfraRuns() int32 {
	if m != nil {
		return m.TotalNonInfraRuns
	}
	return 0
}

func (m *TestInfo) GetPassedNonInfraRuns() int32 {
	if m != nil {
		return m.PassedNonInfraRuns
	}
	return 0
}

func (m *TestInfo) GetFailedNonInfraRuns() int32 {
	if m != nil {
		return m.FailedNonInfraRuns
	}
	return 0
}

func (m *TestInfo) GetFailedInfraRuns() int32 {
	if m != nil {
		return m.FailedInfraRuns
	}
	return 0
}

func (m *TestInfo) GetTotalRunsWithInfra() int32 {
	if m != nil {
		return m.TotalRunsWithInfra
	}
	return 0
}

func (m *TestInfo) GetOtherRuns() int32 {
	if m != nil {
		return m.OtherRuns
	}
	return 0
}

func (m *TestInfo) GetFlakiness() float32 {
	if m != nil {
		return m.Flakiness
	}
	return 0
}

func (m *TestInfo) GetPreviousFlakiness() []float32 {
	if m != nil {
		return m.PreviousFlakiness
	}
	return nil
}

func (m *TestInfo) GetChangeFromLastInterval() TestInfo_Trend {
	if m != nil {
		return m.ChangeFromLastInterval
	}
	return TestInfo_UNKNOWN
}

func (m *TestInfo) GetInfraFailures() map[string]int32 {
	if m != nil {
		return m.InfraFailures
	}
	return nil
}

// Summary of the flakiness and overall healthiness of a dashboard tab
type HealthinessInfo struct {
	// The start of the time frame that the analysis was run for.
	// Represents the lower bound but does not guarantee that the earliest
	// test occurred at start
	Start *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// The end of the time frame that the analysis was run for.
	// Same caveat as above but for upper bound.
	End *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	// A list of test entries associated with this tab + timeframe.
	Tests []*TestInfo `protobuf:"bytes,3,rep,name=tests,proto3" json:"tests,omitempty"`
	// The flakiness out of 100 (think percentage but drop the sign)
	AverageFlakiness float32 `protobuf:"fixed32,4,opt,name=average_flakiness,json=averageFlakiness,proto3" json:"average_flakiness,omitempty"`
	// The average flakiness for previous intervals
	PreviousFlakiness    []float32 `protobuf:"fixed32,5,rep,packed,name=previous_flakiness,json=previousFlakiness,proto3" json:"previous_flakiness,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *HealthinessInfo) Reset()         { *m = HealthinessInfo{} }
func (m *HealthinessInfo) String() string { return proto.CompactTextString(m) }
func (*HealthinessInfo) ProtoMessage()    {}
func (*HealthinessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{2}
}

func (m *HealthinessInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthinessInfo.Unmarshal(m, b)
}
func (m *HealthinessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthinessInfo.Marshal(b, m, deterministic)
}
func (m *HealthinessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthinessInfo.Merge(m, src)
}
func (m *HealthinessInfo) XXX_Size() int {
	return xxx_messageInfo_HealthinessInfo.Size(m)
}
func (m *HealthinessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthinessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HealthinessInfo proto.InternalMessageInfo

func (m *HealthinessInfo) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *HealthinessInfo) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *HealthinessInfo) GetTests() []*TestInfo {
	if m != nil {
		return m.Tests
	}
	return nil
}

func (m *HealthinessInfo) GetAverageFlakiness() float32 {
	if m != nil {
		return m.AverageFlakiness
	}
	return 0
}

func (m *HealthinessInfo) GetPreviousFlakiness() []float32 {
	if m != nil {
		return m.PreviousFlakiness
	}
	return nil
}

// Information about alerts that have been sent
type AlertingData struct {
	// Seconds since epoch at which an email was last sent
	LastEmailTime        *timestamp.Timestamp `protobuf:"bytes,1,opt,name=last_email_time,json=lastEmailTime,proto3" json:"last_email_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AlertingData) Reset()         { *m = AlertingData{} }
func (m *AlertingData) String() string { return proto.CompactTextString(m) }
func (*AlertingData) ProtoMessage()    {}
func (*AlertingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{3}
}

func (m *AlertingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertingData.Unmarshal(m, b)
}
func (m *AlertingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertingData.Marshal(b, m, deterministic)
}
func (m *AlertingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertingData.Merge(m, src)
}
func (m *AlertingData) XXX_Size() int {
	return xxx_messageInfo_AlertingData.Size(m)
}
func (m *AlertingData) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertingData.DiscardUnknown(m)
}

var xxx_messageInfo_AlertingData proto.InternalMessageInfo

func (m *AlertingData) GetLastEmailTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastEmailTime
	}
	return nil
}

// Summary of a dashboard tab.
type DashboardTabSummary struct {
	// The name of the dashboard.
	DashboardName string `protobuf:"bytes,1,opt,name=dashboard_name,json=dashboardName,proto3" json:"dashboard_name,omitempty"`
	// The name of the dashboard tab.
	DashboardTabName string `protobuf:"bytes,2,opt,name=dashboard_tab_name,json=dashboardTabName,proto3" json:"dashboard_tab_name,omitempty"`
	// Any top-level alert on this dashboard tab.
	Alert string `protobuf:"bytes,3,opt,name=alert,proto3" json:"alert,omitempty"`
	// List of failing test summary information.
	FailingTestSummaries []*FailingTestSummary `protobuf:"bytes,4,rep,name=failing_test_summaries,json=failingTestSummaries,proto3" json:"failing_test_summaries,omitempty"`
	// Seconds since epoch at which the test group was last updated.
	LastUpdateTimestamp float64 `protobuf:"fixed64,5,opt,name=last_update_timestamp,json=lastUpdateTimestamp,proto3" json:"last_update_timestamp,omitempty"`
	// A summary of the status of this dashboard tab.
	Status string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	// The overall status for this dashboard tab.
	OverallStatus DashboardTabSummary_TabStatus `protobuf:"varint,7,opt,name=overall_status,json=overallStatus,proto3,enum=DashboardTabSummary_TabStatus" json:"overall_status,omitempty"`
	// The ID for the latest passing build.
	LatestGreen string `protobuf:"bytes,8,opt,name=latest_green,json=latestGreen,proto3" json:"latest_green,omitempty"`
	// Seconds since epoch at which tests last ran.
	LastRunTimestamp float64 `protobuf:"fixed64,9,opt,name=last_run_timestamp,json=lastRunTimestamp,proto3" json:"last_run_timestamp,omitempty"`
	// String indicating the URL for linking to a bug.
	BugUrl string `protobuf:"bytes,10,opt,name=bug_url,json=bugUrl,proto3" json:"bug_url,omitempty"`
	// Metrics for the recent healthiness of a tab
	Healthiness *HealthinessInfo `protobuf:"bytes,12,opt,name=healthiness,proto3" json:"healthiness,omitempty"`
	// All the issue IDs linked to this tab.
	LinkedIssues []string `protobuf:"bytes,13,rep,name=linked_issues,json=linkedIssues,proto3" json:"linked_issues,omitempty"`
	// Metrics about alerts sent with respect to this summary
	// Maintained by alerter; does not need to be populated by summarizer
	AlertingData         *AlertingData `protobuf:"bytes,14,opt,name=alerting_data,json=alertingData,proto3" json:"alerting_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DashboardTabSummary) Reset()         { *m = DashboardTabSummary{} }
func (m *DashboardTabSummary) String() string { return proto.CompactTextString(m) }
func (*DashboardTabSummary) ProtoMessage()    {}
func (*DashboardTabSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{4}
}

func (m *DashboardTabSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashboardTabSummary.Unmarshal(m, b)
}
func (m *DashboardTabSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashboardTabSummary.Marshal(b, m, deterministic)
}
func (m *DashboardTabSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardTabSummary.Merge(m, src)
}
func (m *DashboardTabSummary) XXX_Size() int {
	return xxx_messageInfo_DashboardTabSummary.Size(m)
}
func (m *DashboardTabSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardTabSummary.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardTabSummary proto.InternalMessageInfo

func (m *DashboardTabSummary) GetDashboardName() string {
	if m != nil {
		return m.DashboardName
	}
	return ""
}

func (m *DashboardTabSummary) GetDashboardTabName() string {
	if m != nil {
		return m.DashboardTabName
	}
	return ""
}

func (m *DashboardTabSummary) GetAlert() string {
	if m != nil {
		return m.Alert
	}
	return ""
}

func (m *DashboardTabSummary) GetFailingTestSummaries() []*FailingTestSummary {
	if m != nil {
		return m.FailingTestSummaries
	}
	return nil
}

func (m *DashboardTabSummary) GetLastUpdateTimestamp() float64 {
	if m != nil {
		return m.LastUpdateTimestamp
	}
	return 0
}

func (m *DashboardTabSummary) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DashboardTabSummary) GetOverallStatus() DashboardTabSummary_TabStatus {
	if m != nil {
		return m.OverallStatus
	}
	return DashboardTabSummary_NOT_SET
}

func (m *DashboardTabSummary) GetLatestGreen() string {
	if m != nil {
		return m.LatestGreen
	}
	return ""
}

func (m *DashboardTabSummary) GetLastRunTimestamp() float64 {
	if m != nil {
		return m.LastRunTimestamp
	}
	return 0
}

func (m *DashboardTabSummary) GetBugUrl() string {
	if m != nil {
		return m.BugUrl
	}
	return ""
}

func (m *DashboardTabSummary) GetHealthiness() *HealthinessInfo {
	if m != nil {
		return m.Healthiness
	}
	return nil
}

func (m *DashboardTabSummary) GetLinkedIssues() []string {
	if m != nil {
		return m.LinkedIssues
	}
	return nil
}

func (m *DashboardTabSummary) GetAlertingData() *AlertingData {
	if m != nil {
		return m.AlertingData
	}
	return nil
}

// Summary state of a dashboard.
type DashboardSummary struct {
	// Summary of a dashboard tab; see config.proto.
	TabSummaries         []*DashboardTabSummary `protobuf:"bytes,1,rep,name=tab_summaries,json=tabSummaries,proto3" json:"tab_summaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DashboardSummary) Reset()         { *m = DashboardSummary{} }
func (m *DashboardSummary) String() string { return proto.CompactTextString(m) }
func (*DashboardSummary) ProtoMessage()    {}
func (*DashboardSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{5}
}

func (m *DashboardSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashboardSummary.Unmarshal(m, b)
}
func (m *DashboardSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashboardSummary.Marshal(b, m, deterministic)
}
func (m *DashboardSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardSummary.Merge(m, src)
}
func (m *DashboardSummary) XXX_Size() int {
	return xxx_messageInfo_DashboardSummary.Size(m)
}
func (m *DashboardSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardSummary.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardSummary proto.InternalMessageInfo

func (m *DashboardSummary) GetTabSummaries() []*DashboardTabSummary {
	if m != nil {
		return m.TabSummaries
	}
	return nil
}

func init() {
	proto.RegisterEnum("TestInfo_Trend", TestInfo_Trend_name, TestInfo_Trend_value)
	proto.RegisterEnum("DashboardTabSummary_TabStatus", DashboardTabSummary_TabStatus_name, DashboardTabSummary_TabStatus_value)
	proto.RegisterType((*FailingTestSummary)(nil), "FailingTestSummary")
	proto.RegisterMapType((map[string]string)(nil), "FailingTestSummary.PropertiesEntry")
	proto.RegisterType((*TestInfo)(nil), "TestInfo")
	proto.RegisterMapType((map[string]int32)(nil), "TestInfo.InfraFailuresEntry")
	proto.RegisterType((*HealthinessInfo)(nil), "HealthinessInfo")
	proto.RegisterType((*AlertingData)(nil), "AlertingData")
	proto.RegisterType((*DashboardTabSummary)(nil), "DashboardTabSummary")
	proto.RegisterType((*DashboardSummary)(nil), "DashboardSummary")
}

func init() { proto.RegisterFile("summary.proto", fileDescriptor_f7168d0e3f3f5589) }

var fileDescriptor_f7168d0e3f3f5589 = []byte{
	// 1206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xef, 0x72, 0xdb, 0x44,
	0x10, 0xaf, 0x63, 0xcb, 0x89, 0xd6, 0x96, 0xad, 0x5c, 0x43, 0x11, 0xa1, 0xd0, 0xe0, 0x52, 0xc8,
	0x40, 0x51, 0xc0, 0x0c, 0x33, 0xc0, 0x0c, 0x33, 0x38, 0xa9, 0xdd, 0xba, 0x4d, 0x9d, 0x8e, 0xe2,
	0x4c, 0x87, 0xe1, 0x83, 0xe6, 0x5c, 0x9d, 0x6d, 0x4d, 0x64, 0xc9, 0xa3, 0x3b, 0x85, 0xe6, 0x4d,
	0xe0, 0x4d, 0x78, 0x2a, 0x9e, 0x81, 0xd9, 0x3d, 0xd9, 0x52, 0xdd, 0x40, 0xfb, 0xed, 0xee, 0xb7,
	0xbf, 0xdd, 0xdb, 0xdb, 0xbf, 0x60, 0xc9, 0x6c, 0xb1, 0xe0, 0xe9, 0xb5, 0xbb, 0x4c, 0x13, 0x95,
	0xec, 0xdf, 0x9b, 0x25, 0xc9, 0x2c, 0x12, 0x47, 0x74, 0x9b, 0x64, 0xd3, 0x23, 0x15, 0x2e, 0x84,
	0x54, 0x7c, 0xb1, 0xd4, 0x84, 0xce, 0x9f, 0x75, 0x60, 0x03, 0x1e, 0x46, 0x61, 0x3c, 0x1b, 0x0b,
	0xa9, 0xce, 0xb5, 0x36, 0xfb, 0x0c, 0x9a, 0x41, 0x28, 0x97, 0x11, 0xbf, 0xf6, 0x63, 0xbe, 0x10,
	0x4e, 0xe5, 0xa0, 0x72, 0x68, 0x7a, 0x8d, 0x1c, 0x1b, 0xf1, 0x85, 0x60, 0x1f, 0x83, 0xa9, 0x84,
	0x54, 0x5a, 0xbe, 0x45, 0xf2, 0x1d, 0x04, 0x48, 0xd8, 0x01, 0x6b, 0xca, 0xc3, 0xc8, 0x9f, 0x64,
	0x61, 0x14, 0xf8, 0x61, 0xe0, 0x54, 0xb5, 0x01, 0x04, 0x8f, 0x11, 0x1b, 0x06, 0xec, 0x01, 0xb4,
	0x88, 0xb3, 0x76, 0xc9, 0xa9, 0x1d, 0x54, 0x0e, 0x2b, 0x1e, 0x69, 0x8e, 0x57, 0x20, 0x9a, 0x5a,
	0x72, 0x29, 0x0b, 0x53, 0x86, 0x36, 0x85, 0x60, 0xc9, 0x14, 0x71, 0x0a, 0x53, 0x75, 0x6d, 0x0a,
	0xd1, 0xc2, 0xd4, 0x27, 0x00, 0xf4, 0xe2, 0xab, 0x24, 0x8b, 0x95, 0xb3, 0x7d, 0x50, 0x39, 0x34,
	0x3c, 0x13, 0x91, 0x13, 0x04, 0x50, 0xac, 0x1f, 0x89, 0xc2, 0xf8, 0xd2, 0xd9, 0xa1, 0x67, 0x4c,
	0x42, 0x4e, 0xc3, 0xf8, 0x92, 0x7d, 0x01, 0xed, 0x42, 0xec, 0x2b, 0xf1, 0x5a, 0x39, 0x26, 0x71,
	0xac, 0x35, 0x67, 0x2c, 0x5e, 0x2b, 0xf6, 0x39, 0xb4, 0x34, 0x2f, 0x4b, 0x23, 0x4d, 0x03, 0xa2,
	0x35, 0x09, 0xbd, 0x48, 0x23, 0x62, 0x7d, 0x09, 0x6d, 0x7c, 0x39, 0x4b, 0x85, 0xbf, 0x10, 0x52,
	0xf2, 0x99, 0x70, 0x1a, 0x44, 0x6b, 0xe5, 0xf0, 0x73, 0x8d, 0xb2, 0x7b, 0xd0, 0xc0, 0x07, 0x45,
	0xe0, 0x4f, 0xb2, 0x99, 0x74, 0x9a, 0x07, 0xd5, 0x43, 0xd3, 0x03, 0x0d, 0x1d, 0x67, 0x33, 0x89,
	0xef, 0xe9, 0x38, 0x62, 0x36, 0xc8, 0x75, 0x4b, 0xbf, 0x47, 0x71, 0x14, 0x52, 0x91, 0xf7, 0xdf,
	0xc1, 0x07, 0x11, 0x27, 0xca, 0x06, 0x79, 0x97, 0xc8, 0x4c, 0x0b, 0x07, 0x65, 0x95, 0x23, 0xd8,
	0x2b, 0xab, 0xac, 0x13, 0xd0, 0x22, 0x8d, 0xdd, 0x42, 0x63, 0x95, 0x86, 0x13, 0x80, 0x65, 0x9a,
	0x2c, 0x45, 0xaa, 0x42, 0x21, 0x9d, 0xf6, 0x41, 0xf5, 0xb0, 0xd1, 0xbd, 0xef, 0xbe, 0x5d, 0x5e,
	0xee, 0x8b, 0x35, 0xab, 0x1f, 0xab, 0xf4, 0xda, 0x2b, 0xa9, 0xe1, 0x7f, 0xe7, 0x89, 0x8a, 0x42,
	0xa9, 0xfc, 0x30, 0x90, 0x8e, 0xad, 0xff, 0x9b, 0x43, 0xc3, 0x40, 0x62, 0xe4, 0xc4, 0x02, 0x1d,
	0xe2, 0x41, 0x90, 0x0a, 0x29, 0x85, 0x74, 0x18, 0x91, 0x5a, 0x04, 0xf7, 0x56, 0xe8, 0xfe, 0x2f,
	0xd0, 0xde, 0x78, 0x88, 0xd9, 0x50, 0xbd, 0x14, 0xd7, 0x79, 0x39, 0xe3, 0x91, 0xed, 0x81, 0x71,
	0xc5, 0xa3, 0x6c, 0x55, 0xc2, 0xfa, 0xf2, 0xf3, 0xd6, 0x8f, 0x95, 0xce, 0x5f, 0x06, 0xec, 0xa0,
	0xd3, 0xc3, 0x78, 0x9a, 0xbc, 0x4f, 0x43, 0x1c, 0xc1, 0x9e, 0x4a, 0x14, 0x8f, 0xfc, 0x38, 0x89,
	0xfd, 0x30, 0x9e, 0xa6, 0xdc, 0x4f, 0xb3, 0x58, 0x92, 0x61, 0xc3, 0xdb, 0x25, 0xd9, 0x28, 0x89,
	0x87, 0x28, 0xf1, 0xb2, 0x58, 0x62, 0x4a, 0xb0, 0x3e, 0x45, 0xb0, 0xa9, 0x51, 0x25, 0x0d, 0xa6,
	0x85, 0x9b, 0x2a, 0x98, 0x8b, 0xb7, 0x55, 0x6a, 0x5a, 0x45, 0x0b, 0xdf, 0x50, 0xf9, 0x0a, 0x76,
	0x73, 0x95, 0x12, 0xdd, 0x20, 0x7a, 0x5b, 0x0b, 0xde, 0x30, 0xaf, 0xbf, 0x80, 0x24, 0xff, 0x8f,
	0x50, 0xcd, 0xb5, 0x12, 0xb5, 0x93, 0xe1, 0x31, 0x12, 0x22, 0xf3, 0x65, 0xa8, 0xe6, 0xa4, 0x86,
	0x4d, 0x93, 0xa8, 0xb9, 0x48, 0xb5, 0xdd, 0xbc, 0xa7, 0x08, 0x21, 0x8b, 0x77, 0xc1, 0x9c, 0x46,
	0xfc, 0x32, 0x8c, 0x85, 0x94, 0xd4, 0x52, 0x5b, 0x5e, 0x01, 0xb0, 0x6f, 0x80, 0x2d, 0x53, 0x71,
	0x15, 0x26, 0x99, 0xf4, 0x0b, 0x1a, 0x1c, 0x54, 0x0f, 0xb7, 0xbc, 0xdd, 0x95, 0x64, 0xb0, 0xa6,
	0x3f, 0x85, 0x8f, 0x5e, 0xcd, 0x79, 0x3c, 0x13, 0xfe, 0x34, 0x4d, 0x16, 0x7e, 0xc4, 0xb1, 0x46,
	0x62, 0x25, 0xd2, 0x2b, 0x1e, 0x51, 0x2f, 0xb6, 0xba, 0x6d, 0x77, 0x95, 0x32, 0x77, 0x9c, 0x8a,
	0x38, 0xf0, 0xee, 0x68, 0x8d, 0x41, 0x9a, 0x2c, 0x4e, 0x39, 0x4a, 0x34, 0x9d, 0x9d, 0x40, 0x4b,
	0xc7, 0x23, 0x6f, 0x37, 0xe9, 0x34, 0xa8, 0x5e, 0xef, 0x16, 0x06, 0xe8, 0x83, 0x83, 0x5c, 0xac,
	0x0b, 0xd5, 0x0a, 0xcb, 0xd8, 0xfe, 0xaf, 0xc0, 0xde, 0x26, 0xbd, 0xab, 0xc8, 0x8c, 0x72, 0x91,
	0xfd, 0x00, 0x06, 0xf9, 0xc9, 0x1a, 0xb0, 0x7d, 0x31, 0x7a, 0x36, 0x3a, 0x7b, 0x39, 0xb2, 0x6f,
	0x31, 0x0b, 0xcc, 0xd1, 0x99, 0x7f, 0xf2, 0xa4, 0x37, 0x7a, 0xdc, 0xb7, 0x2b, 0xac, 0x0e, 0x5b,
	0x17, 0x2f, 0xec, 0x2d, 0xb6, 0x03, 0xb5, 0x47, 0x48, 0xa8, 0x76, 0xfe, 0xa9, 0x40, 0xfb, 0x89,
	0xe0, 0x91, 0x9a, 0x53, 0x64, 0xa8, 0x44, 0xbf, 0x05, 0x43, 0x2a, 0x9e, 0x2a, 0x7a, 0xb8, 0xd1,
	0xdd, 0x77, 0xf5, 0xec, 0x77, 0x57, 0xb3, 0xdf, 0x5d, 0x0f, 0x42, 0x4f, 0x13, 0xd9, 0x43, 0xa8,
	0x8a, 0x38, 0x20, 0xa7, 0xfe, 0x9f, 0x8f, 0x34, 0x76, 0x0f, 0x0c, 0x6c, 0x78, 0x2c, 0x4f, 0x0c,
	0x94, 0xb9, 0x0e, 0x94, 0xa7, 0x71, 0xf6, 0x35, 0xec, 0xf2, 0x2b, 0x91, 0x72, 0xcc, 0xcf, 0x3a,
	0x99, 0x35, 0xca, 0xb9, 0x9d, 0x0b, 0x06, 0xef, 0x48, 0xbd, 0xf1, 0x1f, 0xa9, 0xef, 0x78, 0xd0,
	0xec, 0x45, 0xd8, 0xc9, 0xf1, 0xec, 0x11, 0x57, 0x9c, 0x1d, 0x43, 0x9b, 0xd2, 0xaf, 0x27, 0x01,
	0xce, 0xfd, 0xf7, 0xf8, 0xb6, 0x85, 0x2a, 0xfd, 0x45, 0xbe, 0x5e, 0x3a, 0x7f, 0x1b, 0x70, 0xfb,
	0x11, 0x97, 0xf3, 0x49, 0xc2, 0xd3, 0x60, 0xcc, 0x27, 0xab, 0xe5, 0xf7, 0x00, 0x5a, 0xc1, 0x0a,
	0x2e, 0x77, 0xbb, 0xb5, 0x46, 0xa9, 0xdf, 0x1f, 0x02, 0x2b, 0x68, 0x8a, 0x4f, 0xca, 0x9b, 0xd0,
	0x0e, 0x4a, 0x76, 0x89, 0xbd, 0x07, 0x06, 0xc7, 0x0f, 0xe4, 0x9b, 0x50, 0x5f, 0xd8, 0x10, 0xee,
	0x4c, 0xf5, 0x78, 0xd4, 0x13, 0x59, 0x6f, 0x6f, 0x9c, 0x9e, 0x35, 0x0a, 0xf2, 0xed, 0x1b, 0xa6,
	0xa7, 0xb7, 0x37, 0xdd, 0xc4, 0x70, 0x6e, 0x76, 0x71, 0xc0, 0x4b, 0xe5, 0x67, 0xcb, 0x80, 0x2b,
	0x51, 0x5a, 0x85, 0x06, 0xad, 0xc2, 0xdb, 0x28, 0xbc, 0x20, 0x59, 0xb1, 0x10, 0xef, 0x40, 0x5d,
	0x2a, 0xae, 0x32, 0x49, 0x0d, 0x6e, 0x7a, 0xf9, 0x8d, 0xf5, 0xa1, 0x95, 0x60, 0xc2, 0xa2, 0xc8,
	0xcf, 0xe5, 0xdb, 0xd4, 0x5d, 0x9f, 0xba, 0x37, 0xc4, 0xcb, 0xc5, 0x23, 0xb1, 0x3c, 0x2b, 0xd7,
	0xd2, 0x57, 0x1c, 0x9a, 0xf9, 0x02, 0x99, 0xa5, 0x42, 0xc4, 0xf9, 0x4a, 0x6d, 0x68, 0xec, 0x31,
	0x42, 0x18, 0x44, 0xf2, 0x3a, 0xcd, 0xe2, 0x92, 0xcb, 0x26, 0xb9, 0x6c, 0xa3, 0xc4, 0xcb, 0xe2,
	0xc2, 0xdf, 0x0f, 0x61, 0x7b, 0x92, 0xcd, 0x70, 0xb1, 0xe6, 0x3b, 0xb5, 0x3e, 0xc9, 0x66, 0x17,
	0x69, 0xc4, 0xba, 0xd0, 0x98, 0x17, 0xed, 0xe0, 0x34, 0xa9, 0x14, 0x6c, 0x77, 0xa3, 0x45, 0xbc,
	0x32, 0x89, 0xdd, 0x07, 0x2b, 0x5f, 0xac, 0xa1, 0x94, 0x99, 0x90, 0x8e, 0x45, 0x5b, 0xa4, 0xa9,
	0xc1, 0x21, 0x61, 0xac, 0x0b, 0x16, 0xcf, 0xeb, 0xce, 0x0f, 0xb8, 0xe2, 0xb4, 0xfc, 0x1a, 0x5d,
	0xcb, 0x2d, 0x57, 0xa3, 0xd7, 0xe4, 0xa5, 0x5b, 0xe7, 0x77, 0x30, 0xd7, 0x21, 0xc1, 0xbe, 0x1e,
	0x9d, 0x8d, 0xfd, 0xf3, 0xfe, 0xd8, 0xbe, 0x55, 0x6e, 0xf2, 0x0a, 0x76, 0xf3, 0x8b, 0xde, 0xf9,
	0xb9, 0xee, 0xeb, 0x41, 0x6f, 0x78, 0x6a, 0x57, 0x99, 0x09, 0xc6, 0xe0, 0xb4, 0xf7, 0xec, 0x37,
	0xbb, 0x86, 0xc7, 0xf3, 0x71, 0xef, 0xb4, 0x6f, 0x1b, 0x0c, 0xa0, 0x7e, 0xec, 0x9d, 0x3d, 0xeb,
	0x8f, 0xec, 0xfa, 0xd3, 0xda, 0x4e, 0xc3, 0x6e, 0x76, 0x9e, 0x83, 0xbd, 0xce, 0xc4, 0xaa, 0x6c,
	0x7f, 0x02, 0x0b, 0xab, 0xb0, 0x28, 0xa1, 0x0a, 0x95, 0xd0, 0xde, 0x4d, 0x39, 0xf3, 0x9a, 0x6a,
	0x75, 0x0e, 0x85, 0x9c, 0xd4, 0xa9, 0x59, 0xbe, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0x00, 0x83,
	0x8c, 0x63, 0x3e, 0x0a, 0x00, 0x00,
}
