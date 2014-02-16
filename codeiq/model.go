package codeiq

import (
	"errors"
)

const (
	IndexNo              = iota
	IndexChallengerNo    // 挑戦者No
	IndexChallengeCount  // 挑戦回数
	IndexNickname        // ニックネーム
	IndexGender          // 性別
	IndexAge             // 年齢
	IndexResidence       // 居住地
	IndexChallengedOn    // 挑戦完了日時
	IndexFeedbacked      // 評価フィードバック済
	IndexFeedbackComment // 評価フィードバックコメント
	IndexOverallScore    // 総合評価（5段階）
	IndexTechScore       // 技術力（5段階）
	IndexLogialScore     // 論理力（5段階）
	IndexImaginaryScore  // 想像力（5段階）
	IndexQuestionNo      // 設問番号
	IndexAnswerText      // 解答テキスト
)

var (
	AnswerNotFound = errors.New("Answer not found")
)

type Answer struct {
	No             string
	ChallengerNo   string // 挑戦者No
	ChallengeCount string // 挑戦回数
	Nickname       string // ニックネーム
	Gender         string // 性別
	Age            string // 年齢
	Residence      string // 居住地
	ChallengedOn   string // 挑戦完了日時
	Feedbacked     string // 評価フィードバック済
	// 評価フィードバックコメント
	// 総合評価（5段階）
	// 技術力（5段階）
	// 論理力（5段階）
	// 想像力（5段階）
	// 設問番号
	AnswerText string // 解答テキスト
}

func NewAnswerFromRecord(record []string) (answer *Answer) {
	answer = new(Answer)
	answer.No = record[IndexNo]
	answer.ChallengerNo = record[IndexChallengerNo]
	answer.ChallengeCount = record[IndexChallengeCount]
	answer.Nickname = record[IndexNickname]
	answer.Gender = record[IndexGender]
	answer.Age = record[IndexAge]
	answer.Residence = record[IndexResidence]
	answer.ChallengedOn = record[IndexChallengedOn]
	answer.Feedbacked = record[IndexFeedbacked]
	answer.AnswerText = record[IndexAnswerText]

	return
}

type AnswerRepository interface {
	NoOf(no string) (answer Answer, err error)
	Answers() (answers []Answer, err error)
	AddAnswer(answer Answer) (err error)
	CountAnswers() (length int)
	Purge() (err error)
}

type OnMemoryAnswerRepository struct {
	answers []*Answer
}

func NewOnMemoryAnswerRepository() (this *OnMemoryAnswerRepository) {
	this = new(OnMemoryAnswerRepository)
	return
}

func (this *OnMemoryAnswerRepository) NoOf(no string) (answer *Answer, err error) {
	for _, _answer := range this.answers {
		if _answer.No == no {
			return _answer, err
		}
	}

	err = AnswerNotFound

	return
}

func (this *OnMemoryAnswerRepository) Answers() (answers []*Answer, err error) {
	answers = this.answers
	return
}

func (this *OnMemoryAnswerRepository) AddAnswer(answer *Answer) (err error) {
	this.answers = append(this.answers, answer)
	return
}

func (this *OnMemoryAnswerRepository) CountAnswers() (length int) {
	return len(this.answers)
}

func (this *OnMemoryAnswerRepository) Purge() (err error) {
	this.answers = []*Answer{}
	return
}
