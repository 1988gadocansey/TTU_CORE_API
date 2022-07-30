package models

import (
	"gorm.io/gorm"
)

type QualityAssuranceQuestion struct {
	Base
	gorm.Model
	StudentID                  uint8
	Student                    Student       `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"student"`
	comprehensiveOutline       string        `json:"comprehensive_outline,omitempty"`
	outlineBasedOnSyllabus     string        `json:"outline_based_on_syllabus,omitempty"`
	outlineRecommendedBooks    string        `json:"outline_recommended_books,omitempty"`
	lecturerPersonDetails      string        `json:"lecturer_person_details,omitempty"`
	courseObjectiveSpelt       string        `json:"course_objective_spelt,omitempty"`
	courseMaterialList         string        `json:"course_material_list,omitempty"`
	classStartWeek             string        `json:"class_start_week,omitempty"`
	classMetRegularly          string        `json:"class_met_regularly,omitempty"`
	lecturerPunctual           string        `json:"lecturer_punctual,omitempty"`
	lecturerMissedReason       string        `json:"lecturer_missed_reason,omitempty"`
	lecturerStaysPeriod        string        `json:"lecturer_stays_period,omitempty"`
	demonstrateKnowledge       string        `json:"demonstrate_knowledge,omitempty"`
	wellOrganisedDelivery      string        `json:"well_organised_delivery,omitempty"`
	communicateEffectively     string        `json:"communicate_effectively,omitempty"`
	classTimePromLearn         string        `json:"class_time_prom_learn,omitempty"`
	varyingTeachingMeth        string        `json:"varying_teaching_meth,omitempty"`
	encourageStudParticipation string        `json:"encourage_stud_participation,omitempty"`
	encourageProblemSolving    string        `json:"encourage_problem_solving,omitempty"`
	respondToStudConcerns      string        `json:"respond_to_stud_concerns,omitempty"`
	otherMediaDelivery         string        `json:"other_media_delivery,omitempty"`
	roomForQuestion            string        `json:"room_for_question,omitempty"`
	adequateAssignment         string        `json:"adequate_assignment,omitempty"`
	stateFeedbackTime          string        `json:"state_feedback_time,omitempty"`
	markAssignment             string        `json:"mark_assignment,omitempty"`
	discussInClass             string        `json:"discuss_in_class,omitempty"`
	studProgressConcern        string        `json:"stud_progress_concern,omitempty"`
	studResponsibility         string        `json:"stud_responsibility,omitempty"`
	deadlineAssignment         string        `json:"deadline_assignment,omitempty"`
	discloseMarks              string        `json:"disclose_marks,omitempty"`
	lateSubmissionPolicy       string        `json:"late_submission_policy,omitempty"`
	varietyAssignmentUsed      string        `json:"variety_assignment_used,omitempty"`
	courseObjectiveAchieved    string        `json:"course_objective_achieved,omitempty"`
	expectationsCommunicated   string        `json:"expectations_communicated,omitempty"`
	soldHandout                string        `json:"sold_handout,omitempty"`
	createdFriendlyAtmosphere  string        `json:"created_friendly_atmosphere,omitempty"`
	CourseID                   uint8         `  json:"course_id,omitempty" json:"course_id,omitempty"`
	Course                     MountedCourse `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"course"`
	// for old compatibility lets add the lecturer also else the lecturer is in the mounted courses.
	LecturerID uint8    ` json:"lecturer_id,omitempty" json:"lecturer_id,omitempty"`
	Lecturer   User     `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"lecturer"`
	CalenderID uint8    `json:"calender_id,omitempty" json:"calender_id,omitempty"`
	Calender   Calender `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"calender"`
}
