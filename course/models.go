// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package course

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CourseCertificateDetail struct {
	ID                        int32
	CertificateEnabled        pgtype.Bool
	CertificateExpiryDate     pgtype.Date
	CertificateProviderTypeID pgtype.Text
	CertificateProviderType   pgtype.Int4
	CertificateType           pgtype.Text
}

type CourseCourse struct {
	ID                                pgtype.UUID
	Code                              pgtype.Text
	CourseMode                        pgtype.Int4
	CourseStatusID                    pgtype.Int8
	CourseType                        pgtype.Int4
	CreatedBy                         pgtype.Text
	CreatedDate                       pgtype.Timestamp
	Language                          pgtype.Text
	LearningOutcome                   pgtype.Text
	LongDescription                   pgtype.Text
	ShortDescription                  pgtype.Text
	Title                             pgtype.Text
	UpdatedDate                       pgtype.Timestamp
	CertificateDetailID               pgtype.Int8
	CourseInformationID               pgtype.Int8
	CourseStatisticsID                pgtype.Int8
	EndDate                           pgtype.Text
	StartDate                         pgtype.Text
	CourseProviderID                  pgtype.Text
	TypeID                            pgtype.Int8
	Price                             pgtype.Float8
	Availability                      pgtype.Int8
	Assessment                        pgtype.Int8
	Url                               pgtype.Text
	SourceSystemName                  pgtype.Text
	SourceSystemID                    pgtype.Text
	LearningPartnerType               pgtype.Text
	CourseCodeID                      pgtype.Int4
	ReadableCode                      pgtype.Text
	SidCertificateID                  pgtype.Text
	AdditionalAmountToBePaidToPartner pgtype.Float8
	LpCertificateID                   pgtype.Text
	IsEnrollmentDisabled              pgtype.Bool
	Credit                            pgtype.Float8
	ProgramBy                         pgtype.Text
	SchemaID                          pgtype.Text
	ExternalPayment                   pgtype.Int8
	IsPreAssessment                   pgtype.Bool
	PreAssessmentType                 pgtype.Text
	ProjectName                       pgtype.Text
}

type CourseCourseInformation struct {
	ID        int32
	NosCode   []byte
	QpCode    []byte
	SubSector []byte
	SubCourse []byte
	NsqfLevel pgtype.Int4
	CourseID  pgtype.UUID
	Tags      []byte
	Sector    []byte
}

type CourseCourseOccupation struct {
	ID         int32
	Occupation pgtype.Text
	CourseID   pgtype.UUID
}

type CourseCourseStatistic struct {
	ID              int32
	EnrollmentCount pgtype.Int8
	RatingAverage   pgtype.Float8
	RatingTotal     pgtype.Int4
	CourseID        pgtype.Text
}
