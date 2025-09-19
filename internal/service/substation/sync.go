package substation

import (
	"fmt"
	"noize_metter/internal/entities"
	"time"

	"github.com/simonvetter/modbus"
)

var (
	collectInterval = 5 * time.Second
)

func (s *Service) Run() {
	ticker := time.NewTicker(collectInterval)
	defer ticker.Stop()
	for {
		select {
		case <-s.ctx.Done():
			s.log.Info("CF Modbus service stopped.")
			return
		case <-ticker.C:
			if err := s.WrapIteration(); err != nil {
				s.log.Error("error in modbus iteration", err)
			}
		}
	}
}

func (s *Service) WrapIteration() error {
	var err error
	s.mbClient, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL:     s.conf.CFModbusHost,
		Timeout: 10 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("failed to create modbus client: %w", err)
	}
	if err = s.mbClient.Open(); err != nil {
		return fmt.Errorf("failed to open modbus client: %w", err)
	}
	defer s.mbClient.Close()
	dataD, errD := s.readDiscreteRegisters()
	if errD != nil {
		return errD
	}
	dataI, errI := s.readInputRegisters()
	if errI != nil {
		return errI
	}

	data := entities.ModbusRegistersFomRegisters(dataI, dataD)
	s.items.Add(data)
	return nil
}

func (s *Service) readDiscreteRegisters() (*entities.DiscreteRegister, error) {
	values, err := s.mbClient.ReadDiscreteInputs(0, 106)
	if err != nil {
		return nil, fmt.Errorf("failed to read discrete inputs: %w", err)
	}
	return &entities.DiscreteRegister{
		H1ClosedStatus:                         values[0],
		H1OpenStatus:                           values[1],
		ClosedStatus89:                         values[2],
		OpenStatus89:                           values[3],
		A86T1A86BF52H1LockoutRelayOperated:     values[4],
		XFMRT1TopOilTempAlarm:                  values[5],
		XFMRT1WindingTempAlarm:                 values[6],
		XFMRT1HighPressureAlarm:                values[7],
		XFMRT1HighOilLevelMainAlarm:            values[8],
		Trip52H1Coil1:                          values[9],
		Trip52M1Coil1:                          values[10],
		RelayFailureAlarmToSEL2240:             values[11],
		TripLORA86T1:                           values[12],
		MechanicalProtectionAlarm:              values[13],
		Differential:                           values[14],
		PhaseTimeOverCurrent:                   values[15],
		GroundTimeOverCurrent:                  values[16],
		VoltsPerHz:                             values[17],
		PhaseUnderVoltage:                      values[18],
		PhaseOverVoltage:                       values[19],
		NeutralOverVoltage:                     values[20],
		TimeSynch:                              values[21],
		FuseFail:                               values[22],
		CircuitBreakerNOContact:                values[23],
		APhaseInvolvedInTheFault:               values[24],
		BPhaseInvolvedInTheFault:               values[25],
		CPhaseInvolvedInTheFault:               values[26],
		BreakerTrip:                            values[27],
		LossOfPotential:                        values[28],
		TimeSourceAccuracy:                     values[29],
		SoftwareAlarms:                         values[30],
		HardwareAlarm:                          values[31],
		BreakerH1Close:                         values[32],
		BreakerH1Open:                          values[33],
		L1CloseStatus89:                        values[34],
		L1OpenStatus89:                         values[35],
		LOROperate:                             values[36],
		XFMRWindingTemperature:                 values[37],
		XFMRLowPressure:                        values[38],
		XFMRLowOilLevel:                        values[39],
		TripBreakerH1:                          values[40],
		TripBreakerM1:                          values[41],
		TripB86T1:                              values[42],
		TripB86BFH1:                            values[43],
		RelayFailAlarm:                         values[44],
		UnderFrequencyProtection:               values[45],
		OverFrequencyProtection:                values[46],
		DifferentialProtection:                 values[47],
		HVWindingTimeOCProtection:              values[48],
		MVWindingTimeOCProtection:              values[49],
		VoltsPerHz2:                            values[50],
		UnderVoltageProtection:                 values[51],
		OverVoltageProtection:                  values[52],
		OverVoltageNeutralProtection:           values[53],
		HVMeterTimeBasedOnIRIGBTimeSource:      values[54],
		HVMeterIRIGBInHighAccuracyMode:         values[55],
		HVMeterSoftwareAlarm:                   values[56],
		HVMeterHardwareAlarm:                   values[57],
		HVMeterLDPFD3LeadLag:                   values[58],
		HVMeterMainBoardOutput3:                values[59],
		MVMeterTimeBasedOnIRIGBTimeSource:      values[60],
		MVMeterIRIGBInHighAccuracyMode:         values[61],
		MVMeterSoftwareAlarm:                   values[62],
		MVMeterHardwareAlarm:                   values[63],
		MVMeterLDPFD3LeadLag:                   values[64],
		MVMeterMainBoardOutput3:                values[65],
		CommonFailAlarm:                        values[66],
		ATSInAutoMode:                          values[67],
		ATSConnectedToTransformer:              values[68],
		A86BF52T1LORTroubleAlarm:               values[69],
		A86T1LORTroubleAlarm:                   values[70],
		B86BF52T1LORTroubleAlarm:               values[71],
		B86T1LORTroubleAlarm:                   values[72],
		A86BF52T1LOROperated:                   values[73],
		A86T1LOROperated:                       values[74],
		B86BF52T1LOROperated:                   values[75],
		B86T1LOROperated:                       values[76],
		GE845Normal:                            values[77],
		SEL487ENormal:                          values[78],
		SEL2730MNormal:                         values[79],
		SEL2401Normal:                          values[80],
		SEL735HVMETERNormal:                    values[81],
		SEL735MVMETERNormal:                    values[82],
		InverterOutputNormal:                   values[83],
		HVACFailAlarm:                          values[84],
		Section1BattChgrAlarm:                  values[85],
		Section3BattChgrAlarm:                  values[86],
		L1Closed89:                             values[87],
		L1Open89:                               values[88],
		L1LossOfACAlarm89:                      values[89],
		L1LossOfDCAlarm89:                      values[90],
		H1Closed52:                             values[91],
		H1Open52:                               values[92],
		H1LowSF6PressAlarm52:                   values[93],
		H1LowSF6PressCutout52:                  values[94],
		XFMRT1TopOilTempAlarm70C:               values[95],
		XFMRT1WindingTempAlarm75C:              values[96],
		XFMRT1LowN2CylinderPressAlarm200PSI:    values[97],
		XFMRT1PressReliefAlarmMainTank:         values[98],
		XFMRT1LossOfVoltageMainSupplyAlarm:     values[99],
		XFMRT1LossOfVoltageFanSupplyAlarm:      values[100],
		XFMRT1LossOfVoltageCoolingControlAlarm: values[101],
		T1Closed89:                             values[102],
		T1Open89:                               values[103],
		T1LossOfACAlarm89:                      values[104],
		T1LossOfDCAlarm89:                      values[105],
	}, nil
}

func (s *Service) readInputRegisters() (*entities.InputRegister, error) {
	// intervals modbus do not allow to fetch huge range of data by one request, so we split it to several
	intervals := [][]uint16{
		{0, 122},         // 0-122
		{122, 122},       // 122-244
		{244, 338 - 244}, // 244-337
	}
	values := make([]uint16, 0, 338)
	for _, r := range intervals {
		subValues, err := s.mbClient.ReadRegisters(r[0], r[1], modbus.INPUT_REGISTER)
		if err != nil {
			return nil, fmt.Errorf("failed to read registers: %w", err)
		}
		values = append(values, subValues...)
	}

	return &entities.InputRegister{
		InstantaneousTemp:                        ParseFloat32V(values, 0),
		AverageHumidity:                          ParseFloat32V(values, 2),
		MaximumHumidity:                          ParseFloat32V(values, 4),
		MinimumHumidity:                          ParseFloat32V(values, 6),
		AverageAmbientTemperature:                ParseFloat32V(values, 8),
		MaximumAmbientTemperature:                ParseFloat32V(values, 10),
		MinimumAmbientTemperature:                ParseFloat32V(values, 12),
		PhaseACurrent145kV:                       ParseFloat32V(values, 14),
		PhaseBCurrent145kV:                       ParseFloat32V(values, 16),
		PhaseCCurrent145kV:                       ParseFloat32V(values, 18),
		PhaseACurrent38kV:                        ParseFloat32V(values, 20),
		PhaseBCurrent38kV:                        ParseFloat32V(values, 22),
		PhaseCCurrent38kV:                        ParseFloat32V(values, 24),
		PhaseABVoltage145kV:                      ParseFloat32V(values, 26),
		PhaseBCVoltage145kV:                      ParseFloat32V(values, 28),
		PhaseCAVoltage145kV:                      ParseFloat32V(values, 30),
		RealPower:                                ParseFloat32V(values, 32),
		ReactivePower:                            ParseFloat32V(values, 34),
		ApparentPower:                            ParseFloat32V(values, 36),
		PowerFactor:                              ParseFloat32V(values, 38) / 100,
		Frequency:                                ParseFloat32V(values, 40) / 100,
		SidePhaseAN145kV:                         ParseFloat32V(values, 42),
		SidePhaseBN145kV:                         ParseFloat32V(values, 44),
		SidePhaseCN145kV:                         ParseFloat32V(values, 46),
		SideNeutralCurrent38kV:                   ParseFloat32V(values, 48),
		SidePhaseAVoltage145kV:                   ParseFloat32V(values, 50) / 10,
		SidePhaseBVoltage145kV:                   ParseFloat32V(values, 52) / 10,
		SidePhaseCVoltage145kV:                   ParseFloat32V(values, 54) / 10,
		SidePhaseABVoltage145kV:                  ParseFloat32V(values, 56) / 10,
		SidePhaseBCVoltage145kV:                  ParseFloat32V(values, 58) / 10,
		SidePhaseCAVoltage145kV:                  ParseFloat32V(values, 60) / 10,
		SidePhaseACurrent145kV:                   ParseFloat32V(values, 62),
		SidePhaseBCurrent145kV:                   ParseFloat32V(values, 64),
		SidePhaseCCurrent145kV:                   ParseFloat32V(values, 66),
		SidePhaseACurrent38kV:                    ParseFloat32V(values, 68),
		SidePhaseBCurrent38kV:                    ParseFloat32V(values, 70),
		SidePhaseCCurrent38kV:                    ParseFloat32V(values, 72),
		SideRealPower:                            ParseFloat32V(values, 74) / 10,
		SideReactivePower:                        ParseFloat32V(values, 76) / 10,
		SideApparentPower:                        ParseFloat32V(values, 78) / 10,
		Frequency2:                               ParseFloat32V(values, 80) / 100,
		PhasePowerFactorDisplacement:             ParseFloat32V(values, 82) / 100,
		PhaseARMSCurrent:                         ParseFloat32V(values, 84),
		PhaseBRMSCurrent:                         ParseFloat32V(values, 86),
		PhaseCRMSCurrent:                         ParseFloat32V(values, 88),
		NeutralRMSCurrent:                        ParseFloat32V(values, 90),
		PhaseARMSVoltage:                         ParseFloat32V(values, 92) / 10,
		PhaseBRMSVoltage:                         ParseFloat32V(values, 94) / 10,
		PhaseCRMSVoltage:                         ParseFloat32V(values, 96) / 10,
		PhaseABRMSVoltage:                        ParseFloat32V(values, 98) / 10,
		PhaseBCRMSVoltage:                        ParseFloat32V(values, 100) / 10,
		PhaseCARMSVoltage:                        ParseFloat32V(values, 102) / 10,
		RealPower2:                               ParseFloat32V(values, 104) / 10,
		ApparentPower2:                           ParseFloat32V(values, 106) / 10,
		ReactivePower2:                           ParseFloat32V(values, 108) / 10,
		PhaseARealPower:                          ParseFloat32V(values, 110) / 10,
		PhaseBRealPower:                          ParseFloat32V(values, 112) / 10,
		PhaseCRealPower:                          ParseFloat32V(values, 114) / 10,
		PhaseAApparentPower:                      ParseFloat32V(values, 116) / 10,
		PhaseBApparentPower:                      ParseFloat32V(values, 118) / 10,
		PhaseCApparentPower:                      ParseFloat32V(values, 120) / 10,
		PhaseAReactivePower:                      ParseFloat32V(values, 122) / 10,
		PhaseBReactivePower:                      ParseFloat32V(values, 124) / 10,
		PhaseCReactivePower:                      ParseFloat32V(values, 126) / 10,
		Frequency3:                               ParseFloat32V(values, 128) / 100,
		PhaseACurrentMagnitude:                   ParseFloat32V(values, 130),
		PhaseBCurrentMagnitude:                   ParseFloat32V(values, 132),
		PhaseCCurrentMagnitude:                   ParseFloat32V(values, 134),
		NeutralCurrentMagnitude:                  ParseFloat32V(values, 136),
		PhaseAVoltageMagnitude:                   ParseFloat32V(values, 138) / 10,
		PhaseBVoltageMagnitude:                   ParseFloat32V(values, 140) / 10,
		PhaseCVoltageMagnitude:                   ParseFloat32V(values, 142) / 10,
		PhaseABVoltageMagnitude:                  ParseFloat32V(values, 144) / 10,
		PhaseBCVoltageMagnitude:                  ParseFloat32V(values, 146) / 10,
		PhaseCAVoltageMagnitude:                  ParseFloat32V(values, 148) / 10,
		PhaseACurrentAngle:                       ParseFloat32V(values, 150) / 100,
		PhaseBCurrentAngle:                       ParseFloat32V(values, 152) / 100,
		PhaseCCurrentAngle:                       ParseFloat32V(values, 154) / 100,
		NeutralCurrentAngle:                      ParseFloat32V(values, 156) / 100,
		PhaseAVoltageAngle:                       ParseFloat32V(values, 158) / 100,
		PhaseBVoltageAngle:                       ParseFloat32V(values, 160) / 100,
		PhaseCVoltageAngle:                       ParseFloat32V(values, 162) / 100,
		PhaseABVoltageAngle:                      ParseFloat32V(values, 164) / 100,
		PhaseBCVoltageAngle:                      ParseFloat32V(values, 166) / 100,
		PhaseCAVoltageAngle:                      ParseFloat32V(values, 168) / 100,
		DisplacementPowerFactor:                  ParseFloat32V(values, 170) / 100,
		RealPowerDelivered:                       ParseFloat32V(values, 172),
		RealPowerReceived:                        ParseFloat32V(values, 174),
		ApparentPowerDelivered:                   ParseFloat32V(values, 176),
		ApparentPowerReceived:                    ParseFloat32V(values, 178),
		ReactivePowerDelivered:                   ParseFloat32V(values, 180),
		ReactivePowerReceived:                    ParseFloat32V(values, 182),
		NetRealPowerDeliveredReceived:            ParseFloat32V(values, 184),
		NetReactivePowerDeliveredReceived:        ParseFloat32V(values, 186),
		MVMeterPhaseARMSCurrent:                  ParseFloat32V(values, 188),
		MVMeterPhaseBRMSCurrent:                  ParseFloat32V(values, 190),
		MVMeterPhaseCRMSCurrent:                  ParseFloat32V(values, 192),
		MVMeterNeutralRMSCurrent:                 ParseFloat32V(values, 194),
		MVMeterPhaseARMSVoltage:                  ParseFloat32V(values, 196) / 10,
		MVMeterPhaseBRMSVoltage:                  ParseFloat32V(values, 198) / 10,
		MVMeterPhaseCRMSVoltage:                  ParseFloat32V(values, 200) / 10,
		MVMeterPhaseABRMSVoltage:                 ParseFloat32V(values, 202) / 10,
		MVMeterPhaseBCRMSVoltage:                 ParseFloat32V(values, 204) / 10,
		MVMeterPhaseCARMSVoltage:                 ParseFloat32V(values, 206) / 10,
		MVMeterRealPower:                         ParseFloat32V(values, 208) / 10,
		MVMeterApparentPower:                     ParseFloat32V(values, 210) / 10,
		MVMeterReactivePower:                     ParseFloat32V(values, 212) / 10,
		MVMeterPhaseARealPower:                   ParseFloat32V(values, 214) / 10,
		MVMeterPhaseBRealPower:                   ParseFloat32V(values, 216) / 10,
		MVMeterPhaseCRealPower:                   ParseFloat32V(values, 218) / 10,
		MVMeterPhaseAApparentPower:               ParseFloat32V(values, 220) / 10,
		MVMeterPhaseBApparentPower:               ParseFloat32V(values, 222) / 10,
		MVMeterPhaseCApparentPower:               ParseFloat32V(values, 224) / 10,
		MVMeterPhaseAReactivePower:               ParseFloat32V(values, 226) / 10,
		MVMeterPhaseBReactivePower:               ParseFloat32V(values, 228) / 10,
		MVMeterPhaseCReactivePower:               ParseFloat32V(values, 230) / 10,
		MVMeterFrequency:                         ParseFloat32V(values, 232) / 100,
		MVMeterPhaseACurrentMagnitude:            ParseFloat32V(values, 234),
		MVMeterPhaseBCurrentMagnitude:            ParseFloat32V(values, 236),
		MVMeterPhaseCCurrentMagnitude:            ParseFloat32V(values, 238),
		MVMeterNeutralCurrentMagnitude:           ParseFloat32V(values, 240),
		MVMeterPhaseAVoltageMagnitude:            ParseFloat32V(values, 242) / 10,
		MVMeterPhaseBVoltageMagnitude:            ParseFloat32V(values, 244) / 10,
		MVMeterPhaseCVoltageMagnitude:            ParseFloat32V(values, 246) / 10,
		MVMeterPhaseABVoltageMagnitude:           ParseFloat32V(values, 248) / 10,
		MVMeterPhaseBCVoltageMagnitude:           ParseFloat32V(values, 250) / 10,
		MVMeterPhaseCAVoltageMagnitude:           ParseFloat32V(values, 252) / 10,
		MVMeterPhaseACurrentAngle:                ParseFloat32V(values, 254) / 100,
		MVMeterPhaseBCurrentAngle:                ParseFloat32V(values, 256) / 100,
		MVMeterPhaseCCurrentAngle:                ParseFloat32V(values, 258) / 100,
		MVMeterNeutralCurrentAngle:               ParseFloat32V(values, 260) / 100,
		MVMeterPhaseAVoltageAngle:                ParseFloat32V(values, 262) / 100,
		MVMeterPhaseBVoltageAngle:                ParseFloat32V(values, 264) / 100,
		MVMeterPhaseCVoltageAngle:                ParseFloat32V(values, 266) / 100,
		MVMeterPhaseABVoltageAngle:               ParseFloat32V(values, 268) / 100,
		MVMeterPhaseBCVoltageAngle:               ParseFloat32V(values, 270) / 100,
		MVMeterPhaseCAVoltageAngle:               ParseFloat32V(values, 272) / 100,
		MVMeterDisplacementPowerFactor:           ParseFloat32V(values, 274) / 100,
		HVMeterRealPowerDelivered:                ParseFloat32V(values, 276),
		HVMeterRealPowerReceived:                 ParseFloat32V(values, 278),
		HVMeterApparentPowerDelivered:            ParseFloat32V(values, 280),
		HVMeterApparentPowerReceived:             ParseFloat32V(values, 282),
		HVMeterReactivePowerDelivered:            ParseFloat32V(values, 284),
		HVMeterReactivePowerReceived:             ParseFloat32V(values, 286),
		HVMeterNetRealPowerDeliveredReceived:     ParseFloat32V(values, 288),
		HVMeterNetReactivePowerDeliveredReceived: ParseFloat32V(values, 290),
		ATSXfmrPhaseANVoltage:                    ParseFloat32V(values, 292),
		ATSXfmrPhaseBNVoltage:                    ParseFloat32V(values, 294),
		ATSXfmrPhaseABVoltage:                    ParseFloat32V(values, 296),
		ATSGeneratorPhaseANVoltage:               ParseFloat32V(values, 298),
		ATSGeneratorPhaseBNVoltage:               ParseFloat32V(values, 300),
		ATSGeneratorPhaseABVoltage:               ParseFloat32V(values, 302),
		ATSXfmrFrequency:                         ParseFloat32V(values, 304),
		ATSGeneratorFrequency:                    ParseFloat32V(values, 306),
		NumberOfTransfersForGensetToSupply:       ParseFloat32V(values, 308),
		GeneratorFrequency:                       ParseFloat32V(values, 310),
		GeneratorVoltageL3L1:                     ParseFloat32V(values, 312),
		GeneratorCurrentPhaseL1:                  ParseFloat32V(values, 314),
		GeneratorCurrentPhaseL3:                  ParseFloat32V(values, 316),
		GeneratorVoltagePhaseL1N:                 ParseFloat32V(values, 318),
		GeneratorVoltagePhaseL3N:                 ParseFloat32V(values, 320),
		GeneratorFuelLevel:                       ParseFloat32V(values, 322),
		GeneratorCoolantTemp:                     ParseFloat32V(values, 324),
		GeneratorBatteryVoltage:                  ParseFloat32V(values, 326),
		GeneratorEngineSpeed:                     ParseFloat32V(values, 328),
		GeneratorEngineRunTime:                   ParseFloat32V(values, 330),
		GeneratorStatus:                          ParseFloat32V(values, 332),
		GeneratorSuccessStarts:                   ParseFloat32V(values, 334),
		GeneratorOilPressure:                     ParseFloat32V(values, 336),
	}, nil
}
