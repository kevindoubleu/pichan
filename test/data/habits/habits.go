package habits_test

import (
	"github.com/kevindoubleu/pichan/internal/habits"
	pb "github.com/kevindoubleu/pichan/proto/habits"
)

var DomainScorecard1 = habits.Scorecard{
	Name:        "TestInsert-test-name",
	Connotation: "Neutral",
	Time:        "01:02",
	Order:       1,
}
var ProtoScorecard1 = &pb.Scorecard{
	Name:        "TestInsert-test-name",
	Connotation: pb.Connotation_Neutral,
	Time:        "01:02",
	Order:       1,
}
