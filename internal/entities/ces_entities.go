package entities

import (
	"noize_metter/internal/utils"
	"time"
)

type MegaBox struct {
	TimestampNum int64  `parquet:"name=timestamp_num, type=INT64" json:"timestamp_num"`
	TimestampPQ  string `parquet:"name=timestamp, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"timestamp"`

	MegaBoxID           int64  `parquet:"name=megabox_id, type=INT64" json:"megaboxId"`
	IMDC                int64  `parquet:"name=imdc, type=INT64" json:"imdc"`
	PhaseACurrent       int64  `parquet:"name=phaseA_current, type=INT64" json:"phaseA_current"`
	PhaseBCurrent       int64  `parquet:"name=phaseB_current, type=INT64" json:"phaseB_current"`
	PhaseCCurrent       int64  `parquet:"name=phaseC_current, type=INT64" json:"phaseC_current"`
	PhaseNCurrent       int64  `parquet:"name=phaseN_current, type=INT64" json:"phaseN_current"`
	VectorGrdCurrent    int64  `parquet:"name=vector_grd_current, type=INT64" json:"vector_grd_current"`
	PhaseAVoltage       int64  `parquet:"name=phaseA_voltage, type=INT64" json:"phaseA_voltage"`
	PhaseBVoltage       int64  `parquet:"name=phaseB_voltage, type=INT64" json:"phaseB_voltage"`
	PhaseCVoltage       int64  `parquet:"name=phaseC_voltage, type=INT64" json:"phaseC_voltage"`
	LineABVoltage       int64  `parquet:"name=lineAB_voltage, type=INT64" json:"lineAB_voltage"`
	LineBCVoltage       int64  `parquet:"name=lineBC_voltage, type=INT64" json:"lineBC_voltage"`
	LineCAVoltage       int64  `parquet:"name=lineCA_voltage, type=INT64" json:"lineCA_voltage"`
	Frequency           int64  `parquet:"name=frequency, type=INT64" json:"frequency"`
	PhaseSequence       int64  `parquet:"name=phase_sequence, type=INT64" json:"phase_sequence"`
	HeatCapacity        int64  `parquet:"name=heat_capacity, type=INT64" json:"heat_capacity"`
	NoOfOperations      int64  `parquet:"name=no_of_operations, type=INT64" json:"no_of_operations"`
	ContactWear         int64  `parquet:"name=contact_wear, type=INT64" json:"contact_wear"`
	PhaseAActivePower   int64  `parquet:"name=phaseA_activePower, type=INT64" json:"phaseA_activePower"`
	PhaseBActivePower   int64  `parquet:"name=phaseB_activePower, type=INT64" json:"phaseB_activePower"`
	PhaseCActivePower   int64  `parquet:"name=phaseC_activePower, type=INT64" json:"phaseC_activePower"`
	PhaseAReactivePower int64  `parquet:"name=phaseA_reactivePower, type=INT64" json:"phaseA_reactivePower"`
	PhaseBReactivePower int64  `parquet:"name=phaseB_reactivePower, type=INT64" json:"phaseB_reactivePower"`
	PhaseCReactivePower int64  `parquet:"name=phaseC_reactivePower, type=INT64" json:"phaseC_reactivePower"`
	PhaseAApparentPower int64  `parquet:"name=phaseA_apparentPower, type=INT64" json:"phaseA_apparentPower"`
	PhaseBApparentPower int64  `parquet:"name=phaseB_apparentPower, type=INT64" json:"phaseB_apparentPower"`
	PhaseCApparentPower int64  `parquet:"name=phaseC_apparentPower, type=INT64" json:"phaseC_apparentPower"`
	PhaseAPowerFactor   int64  `parquet:"name=phaseA_powerFactor, type=INT64" json:"phaseA_powerFactor"`
	PhaseBPowerFactor   int64  `parquet:"name=phaseB_powerFactor, type=INT64" json:"phaseB_powerFactor"`
	PhaseCPowerFactor   int64  `parquet:"name=phaseC_powerFactor, type=INT64" json:"phaseC_powerFactor"`
	TotalActivePower    int64  `parquet:"name=total_active_power, type=INT64" json:"total_active_power"`
	TotalReactivePower  int64  `parquet:"name=total_reactive_power, type=INT64" json:"total_reactive_power"`
	TotalApparentPower  int64  `parquet:"name=total_apparent_power, type=INT64" json:"total_apparent_power"`
	TotalPowerFactor    int64  `parquet:"name=total_power_factor, type=INT64" json:"total_power_factor"`
	Error               string `parquet:"name=error, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"error"`
}

func (m MegaBox) GetTimestampNum() int64 {
	return m.TimestampNum
}

type Tank struct {
	TankID                int64   `json:"tankId"`
	IMDC                  int64   `json:"imdc"`
	MegaBoxID             int64   `json:"megaboxId"`
	TankInletTemperature  float64 `json:"tank_inlet_temperature,omitempty"`
	TankOutletTemperature float64 `json:"tank_outlet_temperature,omitempty"`
	LiquidLevel           float64 `json:"liquidLevel,omitempty"`
	SPDU                  struct {
		SpduID   int               `json:"spduId"`
		TankID   int               `json:"tankId"`
		Sections []TankSPDUSection `json:"sections"`
	} `json:"spdu,omitempty"`
	FanVFD struct {
		Type          string  `json:"type"`
		OutputEnabled bool    `json:"outputEnabled"`
		Mode          string  `json:"mode"`
		Speed         int64   `json:"speed"`
		Power         float64 `json:"power"`
		Frequency     float64 `json:"frequency"`
		TankID        int     `json:"tankId"`
		Error         string  `json:"error,omitempty"`
	} `json:"fanvfd,omitempty"`
	PumpVFD struct {
		Type          string  `json:"type"`
		OutputEnabled bool    `json:"outputEnabled"`
		Speed         int64   `json:"speed"`
		Power         float64 `json:"power"`
		Frequency     float64 `json:"frequency"`
		TankID        int     `json:"tankId"`
		Error         string  `json:"error,omitempty"`
	} `json:"pumpvfd,omitempty"`
	Error string `json:"error,omitempty"`
}

type TankSPDUSection struct {
	SectionID        int     `json:"sectionId"`
	SpduID           int     `json:"spduId"`
	BoardErrorStatus bool    `json:"board_errorStatus"`
	BoardTemperature float64 `json:"boardTemperature"`
	BoardVoltage     float64 `json:"boardVoltage"`
	Channels         []struct {
		ChannelID    int     `json:"channelId"`
		SectionID    int     `json:"sectionId"`
		Rms          float64 `json:"rms"`
		State        bool    `json:"state"`
		EnableStatus bool    `json:"enableStatus"`
		TripStatus   bool    `json:"tripStatus"`
	} `json:"channels"`
}

func (t *Tank) GetSPDUSection(j int) TankSPDUSection {
	for i := range t.SPDU.Sections {
		if t.SPDU.Sections[i].SectionID == j {
			return t.SPDU.Sections[i]
		}
	}
	return TankSPDUSection{}
}

type FanVFD struct {
	Type          string  `json:"type"`
	OutputEnabled bool    `json:"outputEnabled"`
	Mode          bool    `json:"mode"`
	Speed         int     `json:"speed"`
	Power         float64 `json:"power"`
	Frequency     float64 `json:"frequency"`
	TankID        int     `json:"tankId"`
	Error         string  `json:"error,omitempty"`
}

type PumpVFD struct {
	Type          string  `json:"type"`
	OutputEnabled bool    `json:"outputEnabled"`
	Speed         int     `json:"speed"`
	Power         float64 `json:"power"`
	Frequency     float64 `json:"frequency"`
	TankID        int     `json:"tankId"`
	Error         string  `json:"error,omitempty"`
}

type CesTank struct {
	TimestampPQ  string `parquet:"name=timestamp, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"timestamp"`
	TimestampNum int64  `parquet:"name=timestamp_num, type=INT64" json:"timestamp_num"`

	MegaBoxID             int64   `parquet:"name=megabox_id, type=INT64" json:"megabox_id"`
	TankID                int64   `parquet:"name=tank_id, type=INT64" json:"tank_id"`
	IMDC                  int64   `parquet:"name=imdc, type=INT64" json:"imdc"`
	Error                 string  `parquet:"name=error, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"error"`
	TankInletTemperature  float64 `parquet:"name=tank_inlet_temperature, type=DOUBLE" json:"tank_inlet_temperature"`
	TankOutletTemperature float64 `parquet:"name=tank_outlet_temperature, type=DOUBLE" json:"tank_outlet_temperature"`
	LiquidLevel           float64 `parquet:"name=liquid_level, type=DOUBLE" json:"liquid_level"`

	SPDUSection1Error       bool    `parquet:"name=spdu_section_1_error, type=BOOLEAN" json:"spdu_section_1_error"`
	SPDUSection1Voltage     float64 `parquet:"name=spdu_section_1_voltage, type=DOUBLE" json:"spdu_section_1_voltage"`
	SPDUSection1Temperature float64 `parquet:"name=spdu_section_1_temperature, type=DOUBLE" json:"spdu_section_1_temperature"`

	SPDUSection2Error       bool    `parquet:"name=spdu_section_2_error, type=BOOLEAN" json:"spdu_section_2_error"`
	SPDUSection2Voltage     float64 `parquet:"name=spdu_section_2_voltage, type=DOUBLE" json:"spdu_section_2_voltage"`
	SPDUSection2Temperature float64 `parquet:"name=spdu_section_2_temperature, type=DOUBLE" json:"spdu_section_2_temperature"`

	SPDUSection3Error       bool    `parquet:"name=spdu_section_3_error, type=BOOLEAN" json:"spdu_section_3_error"`
	SPDUSection3Voltage     float64 `parquet:"name=spdu_section_3_voltage, type=DOUBLE" json:"spdu_section_3_voltage"`
	SPDUSection3Temperature float64 `parquet:"name=spdu_section_3_temperature, type=DOUBLE" json:"spdu_section_3_temperature"`

	SPDUSection4Error       bool    `parquet:"name=spdu_section_4_error, type=BOOLEAN" json:"spdu_section_4_error"`
	SPDUSection4Voltage     float64 `parquet:"name=spdu_section_4_voltage, type=DOUBLE" json:"spdu_section_4_voltage"`
	SPDUSection4Temperature float64 `parquet:"name=spdu_section_4_temperature, type=DOUBLE" json:"spdu_section_4_temperature"`

	SPDUSection5Error       bool    `parquet:"name=spdu_section_5_error, type=BOOLEAN" json:"spdu_section_5_error"`
	SPDUSection5Voltage     float64 `parquet:"name=spdu_section_5_voltage, type=DOUBLE" json:"spdu_section_5_voltage"`
	SPDUSection5Temperature float64 `parquet:"name=spdu_section_5_temperature, type=DOUBLE" json:"spdu_section_5_temperature"`

	SPDUSection6Error       bool    `parquet:"name=spdu_section_6_error, type=BOOLEAN" json:"spdu_section_6_error"`
	SPDUSection6Voltage     float64 `parquet:"name=spdu_section_6_voltage, type=DOUBLE" json:"spdu_section_6_voltage"`
	SPDUSection6Temperature float64 `parquet:"name=spdu_section_6_temperature, type=DOUBLE" json:"spdu_section_6_temperature"`

	SPDUSection7Error       bool    `parquet:"name=spdu_section_7_error, type=BOOLEAN" json:"spdu_section_7_error"`
	SPDUSection7Voltage     float64 `parquet:"name=spdu_section_7_voltage, type=DOUBLE" json:"spdu_section_7_voltage"`
	SPDUSection7Temperature float64 `parquet:"name=spdu_section_7_temperature, type=DOUBLE" json:"spdu_section_7_temperature"`

	SPDUSection8Error       bool    `parquet:"name=spdu_section_8_error, type=BOOLEAN" json:"spdu_section_8_error"`
	SPDUSection8Voltage     float64 `parquet:"name=spdu_section_8_voltage, type=DOUBLE" json:"spdu_section_8_voltage"`
	SPDUSection8Temperature float64 `parquet:"name=spdu_section_8_temperature, type=DOUBLE" json:"spdu_section_8_temperature"`

	FanVFDType             string  `parquet:"name=fan_vfd_type, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"fan_vfd_type"`
	FanVFDError            string  `parquet:"name=fan_vfd_error, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"fan_vfd_error"`
	FanVFDOutputEnabled    bool    `parquet:"name=fan_vfd_output_enabled, type=BOOLEAN" json:"fan_vfd_output_enabled"`
	FanVFDMode             bool    `parquet:"name=fan_vfd_mode, type=BOOLEAN" json:"fan_vfd_mode"`
	FanVFDSpeed            int64   `parquet:"name=fan_vfd_speed, type=INT64" json:"fan_vfd_speed"`
	FanVFDPower            float64 `parquet:"name=fan_vfd_power, type=DOUBLE" json:"fan_vfd_power"`
	FanVFDPumpVFDFrequency float64 `parquet:"name=fan_vfd_pump_vfd_frequency, type=DOUBLE" json:"fan_vfd_pump_vfd_frequency"`

	PumpVFDType          string  `parquet:"name=pump_vfd_type, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"pump_vfd_type"`
	PumpVFDError         string  `parquet:"name=pump_vfd_error, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"pump_vfd_error"`
	PumpVFDOutputEnabled bool    `parquet:"name=pump_vfd_output_enabled, type=BOOLEAN" json:"pump_vfd_output_enabled"`
	PumpVFDSpeed         int64   `parquet:"name=pump_vfd_speed, type=INT64" json:"pump_vfd_speed"`
	PumpVFDPower         float64 `parquet:"name=pump_vfd_power, type=DOUBLE" json:"pump_vfd_power"`
	PumpVFDFrequency     float64 `parquet:"name=pump_vfd_frequency, type=DOUBLE" json:"pump_vfd_frequency"`
}

func (m CesTank) GetTimestampNum() int64 {
	return m.TimestampNum
}

func ConvertAPITanks(syncDate time.Time, tList []Tank) []CesTank {
	result := make([]CesTank, 0, len(tList))
	for i := range tList {
		result = append(result, CesTank{
			TimestampNum: utils.TimeToDayIntNum(syncDate),
			TimestampPQ:  syncDate.Format(time.DateTime),

			MegaBoxID:             tList[i].MegaBoxID,
			TankID:                tList[i].TankID,
			IMDC:                  tList[i].IMDC,
			Error:                 tList[i].Error,
			TankInletTemperature:  tList[i].TankInletTemperature,
			TankOutletTemperature: tList[i].TankOutletTemperature,
			LiquidLevel:           tList[i].LiquidLevel,

			SPDUSection1Error:       tList[i].GetSPDUSection(1).BoardErrorStatus,
			SPDUSection1Voltage:     tList[i].GetSPDUSection(1).BoardVoltage,
			SPDUSection1Temperature: tList[i].GetSPDUSection(1).BoardTemperature,

			SPDUSection2Error:       tList[i].GetSPDUSection(2).BoardErrorStatus,
			SPDUSection2Voltage:     tList[i].GetSPDUSection(2).BoardVoltage,
			SPDUSection2Temperature: tList[i].GetSPDUSection(2).BoardTemperature,

			SPDUSection3Error:       tList[i].GetSPDUSection(3).BoardErrorStatus,
			SPDUSection3Voltage:     tList[i].GetSPDUSection(3).BoardVoltage,
			SPDUSection3Temperature: tList[i].GetSPDUSection(3).BoardTemperature,

			SPDUSection4Error:       tList[i].GetSPDUSection(4).BoardErrorStatus,
			SPDUSection4Voltage:     tList[i].GetSPDUSection(4).BoardVoltage,
			SPDUSection4Temperature: tList[i].GetSPDUSection(4).BoardTemperature,

			SPDUSection5Error:       tList[i].GetSPDUSection(5).BoardErrorStatus,
			SPDUSection5Voltage:     tList[i].GetSPDUSection(5).BoardVoltage,
			SPDUSection5Temperature: tList[i].GetSPDUSection(5).BoardTemperature,

			SPDUSection6Error:       tList[i].GetSPDUSection(6).BoardErrorStatus,
			SPDUSection6Voltage:     tList[i].GetSPDUSection(6).BoardVoltage,
			SPDUSection6Temperature: tList[i].GetSPDUSection(6).BoardTemperature,

			SPDUSection7Error:       tList[i].GetSPDUSection(7).BoardErrorStatus,
			SPDUSection7Voltage:     tList[i].GetSPDUSection(7).BoardVoltage,
			SPDUSection7Temperature: tList[i].GetSPDUSection(7).BoardTemperature,

			SPDUSection8Error:       tList[i].GetSPDUSection(8).BoardErrorStatus,
			SPDUSection8Voltage:     tList[i].GetSPDUSection(8).BoardVoltage,
			SPDUSection8Temperature: tList[i].GetSPDUSection(8).BoardTemperature,

			FanVFDType:             tList[i].FanVFD.Type,
			FanVFDError:            tList[i].FanVFD.Error,
			FanVFDOutputEnabled:    tList[i].FanVFD.OutputEnabled,
			FanVFDMode:             tList[i].FanVFD.Mode == "hand",
			FanVFDSpeed:            tList[i].FanVFD.Speed,
			FanVFDPower:            tList[i].FanVFD.Power,
			FanVFDPumpVFDFrequency: tList[i].FanVFD.Frequency,

			PumpVFDType:          tList[i].PumpVFD.Type,
			PumpVFDError:         tList[i].PumpVFD.Error,
			PumpVFDOutputEnabled: tList[i].PumpVFD.OutputEnabled,
			PumpVFDSpeed:         tList[i].PumpVFD.Speed,
			PumpVFDPower:         tList[i].PumpVFD.Power,
			PumpVFDFrequency:     tList[i].PumpVFD.Frequency,
		})
	}
	return result
}

type CesTanksChannelsV2 struct {
	TimestampPQ  string  `json:"timestamp" parquet:"name=timestamp, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	TimestampNum int64   `json:"timestamp_num" parquet:"name=timestamp_num, type=INT64"`
	TankID       int64   `json:"tank_id" parquet:"name=tank_id, type=INT64"`
	MegaboxID    int64   `json:"megabox_id" parquet:"name=megabox_id, type=INT64"`
	IMDC         int64   `json:"imdc" parquet:"name=imdc, type=INT64"`
	Section      int64   `json:"section" parquet:"name=section, type=INT64"`
	Channel      int64   `json:"channel" parquet:"name=channel, type=INT64"`
	Amperage     float64 `json:"amperage" parquet:"name=amperage, type=DOUBLE"`
}

func (m CesTanksChannelsV2) GetTimestampNum() int64 {
	return m.TimestampNum
}

type CesTanksChannels struct {
	TimestampPQ      string  `json:"timestamp" parquet:"name=timestamp, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	TimestampNum     int64   `json:"timestamp_num" parquet:"name=timestamp_num, type=INT64"`
	TankID           int64   `json:"tank_id" parquet:"name=tank_id, type=INT64"`
	MegaboxID        int64   `json:"megabox_id" parquet:"name=megabox_id, type=INT64"`
	IMDC             int64   `json:"imdc" parquet:"name=imdc, type=INT64"`
	Section1Channel1 float64 `json:"section_1_channel_1" parquet:"name=section_1_channel_1, type=DOUBLE"`
	Section1Channel2 float64 `json:"section_1_channel_2" parquet:"name=section_1_channel_2, type=DOUBLE"`
	Section1Channel3 float64 `json:"section_1_channel_3" parquet:"name=section_1_channel_3, type=DOUBLE"`
	Section1Channel4 float64 `json:"section_1_channel_4" parquet:"name=section_1_channel_4, type=DOUBLE"`
	Section1Channel5 float64 `json:"section_1_channel_5" parquet:"name=section_1_channel_5, type=DOUBLE"`
	Section1Channel6 float64 `json:"section_1_channel_6" parquet:"name=section_1_channel_6, type=DOUBLE"`
	Section2Channel1 float64 `json:"section_2_channel_1" parquet:"name=section_2_channel_1, type=DOUBLE"`
	Section2Channel2 float64 `json:"section_2_channel_2" parquet:"name=section_2_channel_2, type=DOUBLE"`
	Section2Channel3 float64 `json:"section_2_channel_3" parquet:"name=section_2_channel_3, type=DOUBLE"`
	Section2Channel4 float64 `json:"section_2_channel_4" parquet:"name=section_2_channel_4, type=DOUBLE"`
	Section2Channel5 float64 `json:"section_2_channel_5" parquet:"name=section_2_channel_5, type=DOUBLE"`
	Section2Channel6 float64 `json:"section_2_channel_6" parquet:"name=section_2_channel_6, type=DOUBLE"`
	Section3Channel1 float64 `json:"section_3_channel_1" parquet:"name=section_3_channel_1, type=DOUBLE"`
	Section3Channel2 float64 `json:"section_3_channel_2" parquet:"name=section_3_channel_2, type=DOUBLE"`
	Section3Channel3 float64 `json:"section_3_channel_3" parquet:"name=section_3_channel_3, type=DOUBLE"`
	Section3Channel4 float64 `json:"section_3_channel_4" parquet:"name=section_3_channel_4, type=DOUBLE"`
	Section3Channel5 float64 `json:"section_3_channel_5" parquet:"name=section_3_channel_5, type=DOUBLE"`
	Section3Channel6 float64 `json:"section_3_channel_6" parquet:"name=section_3_channel_6, type=DOUBLE"`
	Section4Channel1 float64 `json:"section_4_channel_1" parquet:"name=section_4_channel_1, type=DOUBLE"`
	Section4Channel2 float64 `json:"section_4_channel_2" parquet:"name=section_4_channel_2, type=DOUBLE"`
	Section4Channel3 float64 `json:"section_4_channel_3" parquet:"name=section_4_channel_3, type=DOUBLE"`
	Section4Channel4 float64 `json:"section_4_channel_4" parquet:"name=section_4_channel_4, type=DOUBLE"`
	Section4Channel5 float64 `json:"section_4_channel_5" parquet:"name=section_4_channel_5, type=DOUBLE"`
	Section4Channel6 float64 `json:"section_4_channel_6" parquet:"name=section_4_channel_6, type=DOUBLE"`
	Section5Channel1 float64 `json:"section_5_channel_1" parquet:"name=section_5_channel_1, type=DOUBLE"`
	Section5Channel2 float64 `json:"section_5_channel_2" parquet:"name=section_5_channel_2, type=DOUBLE"`
	Section5Channel3 float64 `json:"section_5_channel_3" parquet:"name=section_5_channel_3, type=DOUBLE"`
	Section5Channel4 float64 `json:"section_5_channel_4" parquet:"name=section_5_channel_4, type=DOUBLE"`
	Section5Channel5 float64 `json:"section_5_channel_5" parquet:"name=section_5_channel_5, type=DOUBLE"`
	Section5Channel6 float64 `json:"section_5_channel_6" parquet:"name=section_5_channel_6, type=DOUBLE"`
	Section6Channel1 float64 `json:"section_6_channel_1" parquet:"name=section_6_channel_1, type=DOUBLE"`
	Section6Channel2 float64 `json:"section_6_channel_2" parquet:"name=section_6_channel_2, type=DOUBLE"`
	Section6Channel3 float64 `json:"section_6_channel_3" parquet:"name=section_6_channel_3, type=DOUBLE"`
	Section6Channel4 float64 `json:"section_6_channel_4" parquet:"name=section_6_channel_4, type=DOUBLE"`
	Section6Channel5 float64 `json:"section_6_channel_5" parquet:"name=section_6_channel_5, type=DOUBLE"`
	Section6Channel6 float64 `json:"section_6_channel_6" parquet:"name=section_6_channel_6, type=DOUBLE"`
	Section7Channel1 float64 `json:"section_7_channel_1" parquet:"name=section_7_channel_1, type=DOUBLE"`
	Section7Channel2 float64 `json:"section_7_channel_2" parquet:"name=section_7_channel_2, type=DOUBLE"`
	Section7Channel3 float64 `json:"section_7_channel_3" parquet:"name=section_7_channel_3, type=DOUBLE"`
	Section7Channel4 float64 `json:"section_7_channel_4" parquet:"name=section_7_channel_4, type=DOUBLE"`
	Section7Channel5 float64 `json:"section_7_channel_5" parquet:"name=section_7_channel_5, type=DOUBLE"`
	Section7Channel6 float64 `json:"section_7_channel_6" parquet:"name=section_7_channel_6, type=DOUBLE"`
	Section8Channel1 float64 `json:"section_8_channel_1" parquet:"name=section_8_channel_1, type=DOUBLE"`
	Section8Channel2 float64 `json:"section_8_channel_2" parquet:"name=section_8_channel_2, type=DOUBLE"`
	Section8Channel3 float64 `json:"section_8_channel_3" parquet:"name=section_8_channel_3, type=DOUBLE"`
	Section8Channel4 float64 `json:"section_8_channel_4" parquet:"name=section_8_channel_4, type=DOUBLE"`
	Section8Channel5 float64 `json:"section_8_channel_5" parquet:"name=section_8_channel_5, type=DOUBLE"`
	Section8Channel6 float64 `json:"section_8_channel_6" parquet:"name=section_8_channel_6, type=DOUBLE"`
}

func (m CesTanksChannels) GetTimestampNum() int64 {
	return m.TimestampNum
}
