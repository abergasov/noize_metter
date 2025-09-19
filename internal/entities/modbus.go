package entities

import (
	"noize_metter/internal/utils"
	"time"
)

type ModbusRegisters struct {
	TimestampPQ  string `parquet:"name=timestamp, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"timestamp"`
	TimestampNum int64  `parquet:"name=timestamp_num, type=INT64" json:"timestamp_num"`

	// source annotation to replace is `parquet:"name=%1%%, type=DOUBLE" json:"%1%%"`
	InstantaneousTemp                        float32 `parquet:"name=t1_xfmr_instantaneous_temp_Celsius_ge845, type=DOUBLE" json:"t1_xfmr_instantaneous_temp_Celsius_ge845"`
	AverageHumidity                          float32 `parquet:"name=t1_xfmr_average_humidity_Percent_ge845, type=DOUBLE" json:"t1_xfmr_average_humidity_Percent_ge845"`
	MaximumHumidity                          float32 `parquet:"name=t1_xfmr_maximum_humidity_Percent_ge845, type=DOUBLE" json:"t1_xfmr_maximum_humidity_Percent_ge845"`
	MinimumHumidity                          float32 `parquet:"name=t1_xfmr_minimum_humidity_Percent_ge845, type=DOUBLE" json:"t1_xfmr_minimum_humidity_Percent_ge845"`
	AverageAmbientTemperature                float32 `parquet:"name=t1_xfmr_average_ambient_temperature_Celsius_ge845, type=DOUBLE" json:"t1_xfmr_average_ambient_temperature_Celsius_ge845"`
	MaximumAmbientTemperature                float32 `parquet:"name=t1_xfmr_maximum_ambient_temperature_Celsius_ge845, type=DOUBLE" json:"t1_xfmr_maximum_ambient_temperature_Celsius_ge845"`
	MinimumAmbientTemperature                float32 `parquet:"name=t1_xfmr_minimum_ambient_temperature_Celsius_ge845, type=DOUBLE" json:"t1_xfmr_minimum_ambient_temperature_Celsius_ge845"`
	PhaseACurrent145kV                       float32 `parquet:"name=t1_xfmr_145kv_side_phase_a_current_A_ge845, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_a_current_A_ge845"`
	PhaseBCurrent145kV                       float32 `parquet:"name=t1_xfmr_145kv_side_phase_b_current_A_ge845, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_b_current_A_ge845"`
	PhaseCCurrent145kV                       float32 `parquet:"name=t1_xfmr_145kv_side_phase_c_current_A_ge845, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_c_current_A_ge845"`
	PhaseACurrent38kV                        float32 `parquet:"name=t1_xfmr_38kv_side_phase_a_current_A_ge845, type=DOUBLE" json:"t1_xfmr_38kv_side_phase_a_current_A_ge845"`
	PhaseBCurrent38kV                        float32 `parquet:"name=t1_xfmr_38kv_side_phase_b_current_A_ge845, type=DOUBLE" json:"t1_xfmr_38kv_side_phase_b_current_A_ge845"`
	PhaseCCurrent38kV                        float32 `parquet:"name=t1_xfmr_38kv_side_phase_c_current_A_ge845, type=DOUBLE" json:"t1_xfmr_38kv_side_phase_c_current_A_ge845"`
	PhaseABVoltage145kV                      float32 `parquet:"name=t1_xfmr_145kv_side_phase_ab_voltage_V_ge845, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_ab_voltage_V_ge845"`
	PhaseBCVoltage145kV                      float32 `parquet:"name=t1_xfmr_145kv_side_phase_bc_voltage_V_ge845, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_bc_voltage_V_ge845"`
	PhaseCAVoltage145kV                      float32 `parquet:"name=t1_xfmr_145kv_side_phase_ca_voltage_V_ge845, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_ca_voltage_V_ge845"`
	RealPower                                float32 `parquet:"name=t1_xfmr_real_power_KW_ge845, type=DOUBLE" json:"t1_xfmr_real_power_KW_ge845"`
	ReactivePower                            float32 `parquet:"name=t1_xfmr_reactive_power_KVAR_ge845, type=DOUBLE" json:"t1_xfmr_reactive_power_KVAR_ge845"`
	ApparentPower                            float32 `parquet:"name=t1_xfmr_apparent_power_KVAR_ge845, type=DOUBLE" json:"t1_xfmr_apparent_power_KVAR_ge845"`
	PowerFactor                              float32 `parquet:"name=t1_xfmr_power_factor_ge845, type=DOUBLE" json:"t1_xfmr_power_factor_ge845"`
	Frequency                                float32 `parquet:"name=t1_xfmr_frequency_HZ_ge845, type=DOUBLE" json:"t1_xfmr_frequency_HZ_ge845"`
	SidePhaseAN145kV                         float32 `parquet:"name=t1_xfmr_145kv_side_phase_an_voltage_V_ge845, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_an_voltage_V_ge845"`
	SidePhaseBN145kV                         float32 `parquet:"name=t1_xfmr_145kv_side_phase_bn_voltage_V_ge845, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_bn_voltage_V_ge845"`
	SidePhaseCN145kV                         float32 `parquet:"name=t1_xfmr_145kv_side_phase_cn_voltage_V_ge845, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_cn_voltage_V_ge845"`
	SideNeutralCurrent38kV                   float32 `parquet:"name=t1_xfmr_38kv_side_neutral_current_A_ge845, type=DOUBLE" json:"t1_xfmr_38kv_side_neutral_current_A_ge845"`
	SidePhaseAVoltage145kV                   float32 `parquet:"name=t1_xfmr_145kv_side_phase_a_voltage_Kv_sel487e, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_a_voltage_Kv_sel487e"`
	SidePhaseBVoltage145kV                   float32 `parquet:"name=t1_xfmr_145kv_side_phase_b_voltage_Kv_sel487e, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_b_voltage_Kv_sel487e"`
	SidePhaseCVoltage145kV                   float32 `parquet:"name=t1_xfmr_145kv_side_phase_c_voltage_Kv_sel487e, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_c_voltage_Kv_sel487e"`
	SidePhaseABVoltage145kV                  float32 `parquet:"name=t1_xfmr_145kv_side_phase_ab_voltage_Kv_sel487e, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_ab_voltage_Kv_sel487e"`
	SidePhaseBCVoltage145kV                  float32 `parquet:"name=t1_xfmr_145kv_side_phase_bc_voltage_Kv_sel487e, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_bc_voltage_Kv_sel487e"`
	SidePhaseCAVoltage145kV                  float32 `parquet:"name=t1_xfmr_145kv_side_phase_ca_voltage_Kv_sel487e, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_ca_voltage_Kv_sel487e"`
	SidePhaseACurrent145kV                   float32 `parquet:"name=t1_xfmr_145kv_side_phase_a_current_A_sel487e, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_a_current_A_sel487e"`
	SidePhaseBCurrent145kV                   float32 `parquet:"name=t1_xfmr_145kv_side_phase_b_current_A_sel487e, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_b_current_A_sel487e"`
	SidePhaseCCurrent145kV                   float32 `parquet:"name=t1_xfmr_145kv_side_phase_c_current_A_sel487e, type=DOUBLE" json:"t1_xfmr_145kv_side_phase_c_current_A_sel487e"`
	SidePhaseACurrent38kV                    float32 `parquet:"name=t1_xfmr_38kv_side_phase_a_current_A_sel487e, type=DOUBLE" json:"t1_xfmr_38kv_side_phase_a_current_A_sel487e"`
	SidePhaseBCurrent38kV                    float32 `parquet:"name=t1_xfmr_38kv_side_phase_b_current_A_sel487e, type=DOUBLE" json:"t1_xfmr_38kv_side_phase_b_current_A_sel487e"`
	SidePhaseCCurrent38kV                    float32 `parquet:"name=t1_xfmr_38kv_side_phase_c_current_A_sel487e, type=DOUBLE" json:"t1_xfmr_38kv_side_phase_c_current_A_sel487e"`
	SideRealPower                            float32 `parquet:"name=t1_xfmr_38kv_side_real_power_MW_sel487e, type=DOUBLE" json:"t1_xfmr_38kv_side_real_power_MW_sel487e"`
	SideReactivePower                        float32 `parquet:"name=t1_xfmr_38kv_side_reactive_power_MVAR_sel487e, type=DOUBLE" json:"t1_xfmr_38kv_side_reactive_power_MVAR_sel487e"`
	SideApparentPower                        float32 `parquet:"name=t1_xfmr_38kv_side_apparent_power_MVA_sel487e, type=DOUBLE" json:"t1_xfmr_38kv_side_apparent_power_MVA_sel487e"`
	Frequency2                               float32 `parquet:"name=t1_xfmr_frequency_HZ_sel487e, type=DOUBLE" json:"t1_xfmr_frequency_HZ_sel487e"`
	PhasePowerFactorDisplacement             float32 `parquet:"name=t1_xfmr_3_phase_power_factor_displacement_sel487e, type=DOUBLE" json:"t1_xfmr_3_phase_power_factor_displacement_sel487e"`
	PhaseARMSCurrent                         float32 `parquet:"name=hv_meter_phase_a_rms_current_A_sel735hv, type=DOUBLE" json:"hv_meter_phase_a_rms_current_A_sel735hv"`
	PhaseBRMSCurrent                         float32 `parquet:"name=hv_meter_phase_b_rms_current_A_sel735hv, type=DOUBLE" json:"hv_meter_phase_b_rms_current_A_sel735hv"`
	PhaseCRMSCurrent                         float32 `parquet:"name=hv_meter_phase_c_rms_current_A_sel735hv, type=DOUBLE" json:"hv_meter_phase_c_rms_current_A_sel735hv"`
	NeutralRMSCurrent                        float32 `parquet:"name=hv_meter_neutral_rms_current_A_sel735hv, type=DOUBLE" json:"hv_meter_neutral_rms_current_A_sel735hv"`
	PhaseARMSVoltage                         float32 `parquet:"name=hv_meter_phase_a_rms_voltage_Kv_sel735hv, type=DOUBLE" json:"hv_meter_phase_a_rms_voltage_Kv_sel735hv"`
	PhaseBRMSVoltage                         float32 `parquet:"name=hv_meter_phase_b_rms_voltage_Kv_sel735hv, type=DOUBLE" json:"hv_meter_phase_b_rms_voltage_Kv_sel735hv"`
	PhaseCRMSVoltage                         float32 `parquet:"name=hv_meter_phase_c_rms_voltage_Kv_sel735hv, type=DOUBLE" json:"hv_meter_phase_c_rms_voltage_Kv_sel735hv"`
	PhaseABRMSVoltage                        float32 `parquet:"name=hv_meter_phase_ab_rms_voltage_Kv_sel735hv, type=DOUBLE" json:"hv_meter_phase_ab_rms_voltage_Kv_sel735hv"`
	PhaseBCRMSVoltage                        float32 `parquet:"name=hv_meter_phase_bc_rms_voltage_Kv_sel735hv, type=DOUBLE" json:"hv_meter_phase_bc_rms_voltage_Kv_sel735hv"`
	PhaseCARMSVoltage                        float32 `parquet:"name=hv_meter_phase_ca_rms_voltage_Kv_sel735hv, type=DOUBLE" json:"hv_meter_phase_ca_rms_voltage_Kv_sel735hv"`
	RealPower2                               float32 `parquet:"name=hv_meter_real_power_MW_sel735hv, type=DOUBLE" json:"hv_meter_real_power_MW_sel735hv"`
	ApparentPower2                           float32 `parquet:"name=hv_meter_apparent_power_MVA_sel735hv, type=DOUBLE" json:"hv_meter_apparent_power_MVA_sel735hv"`
	ReactivePower2                           float32 `parquet:"name=hv_meter_reactive_power_MVAR_sel735hv, type=DOUBLE" json:"hv_meter_reactive_power_MVAR_sel735hv"`
	PhaseARealPower                          float32 `parquet:"name=hv_meter_phase_a_real_power_MW_sel735hv, type=DOUBLE" json:"hv_meter_phase_a_real_power_MW_sel735hv"`
	PhaseBRealPower                          float32 `parquet:"name=hv_meter_phase_b_real_power_MW_sel735hv, type=DOUBLE" json:"hv_meter_phase_b_real_power_MW_sel735hv"`
	PhaseCRealPower                          float32 `parquet:"name=hv_meter_phase_c_real_power_MW_sel735hv, type=DOUBLE" json:"hv_meter_phase_c_real_power_MW_sel735hv"`
	PhaseAApparentPower                      float32 `parquet:"name=hv_meter_phase_a_apparent_power_MVA_sel735hv, type=DOUBLE" json:"hv_meter_phase_a_apparent_power_MVA_sel735hv"`
	PhaseBApparentPower                      float32 `parquet:"name=hv_meter_phase_b_apparent_power_MVA_sel735hv, type=DOUBLE" json:"hv_meter_phase_b_apparent_power_MVA_sel735hv"`
	PhaseCApparentPower                      float32 `parquet:"name=hv_meter_phase_c_apparent_power_MVA_sel735hv, type=DOUBLE" json:"hv_meter_phase_c_apparent_power_MVA_sel735hv"`
	PhaseAReactivePower                      float32 `parquet:"name=hv_meter_phase_a_reactive_power_MVAR_sel735hv, type=DOUBLE" json:"hv_meter_phase_a_reactive_power_MVAR_sel735hv"`
	PhaseBReactivePower                      float32 `parquet:"name=hv_meter_phase_b_reactive_power_MVAR_sel735hv, type=DOUBLE" json:"hv_meter_phase_b_reactive_power_MVAR_sel735hv"`
	PhaseCReactivePower                      float32 `parquet:"name=hv_meter_phase_c_reactive_power_MVAR_sel735hv, type=DOUBLE" json:"hv_meter_phase_c_reactive_power_MVAR_sel735hv"`
	Frequency3                               float32 `parquet:"name=hv_meter_frequency_HZ_sel735hv, type=DOUBLE" json:"hv_meter_frequency_HZ_sel735hv"`
	PhaseACurrentMagnitude                   float32 `parquet:"name=hv_meter_phase_a_current_magnitude_A_sel735hv, type=DOUBLE" json:"hv_meter_phase_a_current_magnitude_A_sel735hv"`
	PhaseBCurrentMagnitude                   float32 `parquet:"name=hv_meter_phase_b_current_magnitude_A_sel735hv, type=DOUBLE" json:"hv_meter_phase_b_current_magnitude_A_sel735hv"`
	PhaseCCurrentMagnitude                   float32 `parquet:"name=hv_meter_phase_c_current_magnitude_A_sel735hv, type=DOUBLE" json:"hv_meter_phase_c_current_magnitude_A_sel735hv"`
	NeutralCurrentMagnitude                  float32 `parquet:"name=hv_meter_neutral_current_magnitude_A_sel735hv, type=DOUBLE" json:"hv_meter_neutral_current_magnitude_A_sel735hv"`
	PhaseAVoltageMagnitude                   float32 `parquet:"name=hv_meter_phase_a_voltage_magnitude_Kv_sel735hv, type=DOUBLE" json:"hv_meter_phase_a_voltage_magnitude_Kv_sel735hv"`
	PhaseBVoltageMagnitude                   float32 `parquet:"name=hv_meter_phase_b_voltage_magnitude_Kv_sel735hv, type=DOUBLE" json:"hv_meter_phase_b_voltage_magnitude_Kv_sel735hv"`
	PhaseCVoltageMagnitude                   float32 `parquet:"name=hv_meter_phase_c_voltage_magnitude_Kv_sel735hv, type=DOUBLE" json:"hv_meter_phase_c_voltage_magnitude_Kv_sel735hv"`
	PhaseABVoltageMagnitude                  float32 `parquet:"name=hv_meter_phase_ab_voltage_magnitude_Kv_sel735hv, type=DOUBLE" json:"hv_meter_phase_ab_voltage_magnitude_Kv_sel735hv"`
	PhaseBCVoltageMagnitude                  float32 `parquet:"name=hv_meter_phase_bc_voltage_magnitude_Kv_sel735hv, type=DOUBLE" json:"hv_meter_phase_bc_voltage_magnitude_Kv_sel735hv"`
	PhaseCAVoltageMagnitude                  float32 `parquet:"name=hv_meter_phase_ca_voltage_magnitude_Kv_sel735hv, type=DOUBLE" json:"hv_meter_phase_ca_voltage_magnitude_Kv_sel735hv"`
	PhaseACurrentAngle                       float32 `parquet:"name=hv_meter_phase_a_current_angle_Deg_sel735hv, type=DOUBLE" json:"hv_meter_phase_a_current_angle_Deg_sel735hv"`
	PhaseBCurrentAngle                       float32 `parquet:"name=hv_meter_phase_b_current_angle_Deg_sel735hv, type=DOUBLE" json:"hv_meter_phase_b_current_angle_Deg_sel735hv"`
	PhaseCCurrentAngle                       float32 `parquet:"name=hv_meter_phase_c_current_angle_Deg_sel735hv, type=DOUBLE" json:"hv_meter_phase_c_current_angle_Deg_sel735hv"`
	NeutralCurrentAngle                      float32 `parquet:"name=hv_meter_neutral_current_angle_Deg_sel735hv, type=DOUBLE" json:"hv_meter_neutral_current_angle_Deg_sel735hv"`
	PhaseAVoltageAngle                       float32 `parquet:"name=hv_meter_phase_a_voltage_angle_Deg_sel735hv, type=DOUBLE" json:"hv_meter_phase_a_voltage_angle_Deg_sel735hv"`
	PhaseBVoltageAngle                       float32 `parquet:"name=hv_meter_phase_b_voltage_angle_Deg_sel735hv, type=DOUBLE" json:"hv_meter_phase_b_voltage_angle_Deg_sel735hv"`
	PhaseCVoltageAngle                       float32 `parquet:"name=hv_meter_phase_c_voltage_angle_Deg_sel735hv, type=DOUBLE" json:"hv_meter_phase_c_voltage_angle_Deg_sel735hv"`
	PhaseABVoltageAngle                      float32 `parquet:"name=hv_meter_phase_ab_voltage_angle_Deg_sel735hv, type=DOUBLE" json:"hv_meter_phase_ab_voltage_angle_Deg_sel735hv"`
	PhaseBCVoltageAngle                      float32 `parquet:"name=hv_meter_phase_bc_voltage_angle_Deg_sel735hv, type=DOUBLE" json:"hv_meter_phase_bc_voltage_angle_Deg_sel735hv"`
	PhaseCAVoltageAngle                      float32 `parquet:"name=hv_meter_phase_ca_voltage_angle_Deg_sel735hv, type=DOUBLE" json:"hv_meter_phase_ca_voltage_angle_Deg_sel735hv"`
	DisplacementPowerFactor                  float32 `parquet:"name=hv_meter_displacement_power_factor_sel735hv, type=DOUBLE" json:"hv_meter_displacement_power_factor_sel735hv"`
	RealPowerDelivered                       float32 `parquet:"name=hv_meter_real_power_delivered_KWh_sel735hv, type=DOUBLE" json:"hv_meter_real_power_delivered_KWh_sel735hv"`
	RealPowerReceived                        float32 `parquet:"name=hv_meter_real_power_received_KWh_sel735hv, type=DOUBLE" json:"hv_meter_real_power_received_KWh_sel735hv"`
	ApparentPowerDelivered                   float32 `parquet:"name=hv_meter_apparent_power_delivered_KVAh_sel735hv, type=DOUBLE" json:"hv_meter_apparent_power_delivered_KVAh_sel735hv"`
	ApparentPowerReceived                    float32 `parquet:"name=hv_meter_apparent_power_received_KVAh_sel735hv, type=DOUBLE" json:"hv_meter_apparent_power_received_KVAh_sel735hv"`
	ReactivePowerDelivered                   float32 `parquet:"name=hv_meter_reactive_power_delivered_KVarh_sel735hv, type=DOUBLE" json:"hv_meter_reactive_power_delivered_KVarh_sel735hv"`
	ReactivePowerReceived                    float32 `parquet:"name=hv_meter_reactive_power_received_KVarh_sel735hv, type=DOUBLE" json:"hv_meter_reactive_power_received_KVarh_sel735hv"`
	NetRealPowerDeliveredReceived            float32 `parquet:"name=hv_meter_net_real_power_delivered_received_KWh_sel735hv, type=DOUBLE" json:"hv_meter_net_real_power_delivered_received_KWh_sel735hv"`
	NetReactivePowerDeliveredReceived        float32 `parquet:"name=hv_meter_net_reactive_power_delivered_received_KVarh_sel735hv, type=DOUBLE" json:"hv_meter_net_reactive_power_delivered_received_KVarh_sel735hv"`
	MVMeterPhaseARMSCurrent                  float32 `parquet:"name=mv_meter_phase_a_rms_current_A_sel735mv, type=DOUBLE" json:"mv_meter_phase_a_rms_current_A_sel735mv"`
	MVMeterPhaseBRMSCurrent                  float32 `parquet:"name=mv_meter_phase_b_rms_current_A_sel735mv, type=DOUBLE" json:"mv_meter_phase_b_rms_current_A_sel735mv"`
	MVMeterPhaseCRMSCurrent                  float32 `parquet:"name=mv_meter_phase_c_rms_current_A_sel735mv, type=DOUBLE" json:"mv_meter_phase_c_rms_current_A_sel735mv"`
	MVMeterNeutralRMSCurrent                 float32 `parquet:"name=mv_meter_neutral_rms_current_A_sel735mv, type=DOUBLE" json:"mv_meter_neutral_rms_current_A_sel735mv"`
	MVMeterPhaseARMSVoltage                  float32 `parquet:"name=mv_meter_phase_a_rms_voltage_Kv_sel735mv, type=DOUBLE" json:"mv_meter_phase_a_rms_voltage_Kv_sel735mv"`
	MVMeterPhaseBRMSVoltage                  float32 `parquet:"name=mv_meter_phase_b_rms_voltage_Kv_sel735mv, type=DOUBLE" json:"mv_meter_phase_b_rms_voltage_Kv_sel735mv"`
	MVMeterPhaseCRMSVoltage                  float32 `parquet:"name=mv_meter_phase_c_rms_voltage_Kv_sel735mv, type=DOUBLE" json:"mv_meter_phase_c_rms_voltage_Kv_sel735mv"`
	MVMeterPhaseABRMSVoltage                 float32 `parquet:"name=mv_meter_phase_ab_rms_voltage_Kv_sel735mv, type=DOUBLE" json:"mv_meter_phase_ab_rms_voltage_Kv_sel735mv"`
	MVMeterPhaseBCRMSVoltage                 float32 `parquet:"name=mv_meter_phase_bc_rms_voltage_Kv_sel735mv, type=DOUBLE" json:"mv_meter_phase_bc_rms_voltage_Kv_sel735mv"`
	MVMeterPhaseCARMSVoltage                 float32 `parquet:"name=mv_meter_phase_ca_rms_voltage_Kv_sel735mv, type=DOUBLE" json:"mv_meter_phase_ca_rms_voltage_Kv_sel735mv"`
	MVMeterRealPower                         float32 `parquet:"name=mv_meter_real_power_MW_sel735mv, type=DOUBLE" json:"mv_meter_real_power_MW_sel735mv"`
	MVMeterApparentPower                     float32 `parquet:"name=mv_meter_apparent_power_MVA_sel735mv, type=DOUBLE" json:"mv_meter_apparent_power_MVA_sel735mv"`
	MVMeterReactivePower                     float32 `parquet:"name=mv_meter_reactive_power_MVAR_sel735mv, type=DOUBLE" json:"mv_meter_reactive_power_MVAR_sel735mv"`
	MVMeterPhaseARealPower                   float32 `parquet:"name=mv_meter_phase_a_real_power_MW_sel735mv, type=DOUBLE" json:"mv_meter_phase_a_real_power_MW_sel735mv"`
	MVMeterPhaseBRealPower                   float32 `parquet:"name=mv_meter_phase_b_real_power_MW_sel735mv, type=DOUBLE" json:"mv_meter_phase_b_real_power_MW_sel735mv"`
	MVMeterPhaseCRealPower                   float32 `parquet:"name=mv_meter_phase_c_real_power_MW_sel735mv, type=DOUBLE" json:"mv_meter_phase_c_real_power_MW_sel735mv"`
	MVMeterPhaseAApparentPower               float32 `parquet:"name=mv_meter_phase_a_apparent_power_MVA_sel735mv, type=DOUBLE" json:"mv_meter_phase_a_apparent_power_MVA_sel735mv"`
	MVMeterPhaseBApparentPower               float32 `parquet:"name=mv_meter_phase_b_apparent_power_MVA_sel735mv, type=DOUBLE" json:"mv_meter_phase_b_apparent_power_MVA_sel735mv"`
	MVMeterPhaseCApparentPower               float32 `parquet:"name=mv_meter_phase_c_apparent_power_MVA_sel735mv, type=DOUBLE" json:"mv_meter_phase_c_apparent_power_MVA_sel735mv"`
	MVMeterPhaseAReactivePower               float32 `parquet:"name=mv_meter_phase_a_reactive_power_MVAR_sel735mv, type=DOUBLE" json:"mv_meter_phase_a_reactive_power_MVAR_sel735mv"`
	MVMeterPhaseBReactivePower               float32 `parquet:"name=mv_meter_phase_b_reactive_power_MVAR_sel735mv, type=DOUBLE" json:"mv_meter_phase_b_reactive_power_MVAR_sel735mv"`
	MVMeterPhaseCReactivePower               float32 `parquet:"name=mv_meter_phase_c_reactive_power_MVAR_sel735mv, type=DOUBLE" json:"mv_meter_phase_c_reactive_power_MVAR_sel735mv"`
	MVMeterFrequency                         float32 `parquet:"name=mv_meter_frequency_HZ_sel735mv, type=DOUBLE" json:"mv_meter_frequency_HZ_sel735mv"`
	MVMeterPhaseACurrentMagnitude            float32 `parquet:"name=mv_meter_phase_a_current_magnitude_A_sel735mv, type=DOUBLE" json:"mv_meter_phase_a_current_magnitude_A_sel735mv"`
	MVMeterPhaseBCurrentMagnitude            float32 `parquet:"name=mv_meter_phase_b_current_magnitude_A_sel735mv, type=DOUBLE" json:"mv_meter_phase_b_current_magnitude_A_sel735mv"`
	MVMeterPhaseCCurrentMagnitude            float32 `parquet:"name=mv_meter_phase_c_current_magnitude_A_sel735mv, type=DOUBLE" json:"mv_meter_phase_c_current_magnitude_A_sel735mv"`
	MVMeterNeutralCurrentMagnitude           float32 `parquet:"name=mv_meter_neutral_current_magnitude_A_sel735mv, type=DOUBLE" json:"mv_meter_neutral_current_magnitude_A_sel735mv"`
	MVMeterPhaseAVoltageMagnitude            float32 `parquet:"name=mv_meter_phase_a_voltage_magnitude_Kv_sel735mv, type=DOUBLE" json:"mv_meter_phase_a_voltage_magnitude_Kv_sel735mv"`
	MVMeterPhaseBVoltageMagnitude            float32 `parquet:"name=mv_meter_phase_b_voltage_magnitude_Kv_sel735mv, type=DOUBLE" json:"mv_meter_phase_b_voltage_magnitude_Kv_sel735mv"`
	MVMeterPhaseCVoltageMagnitude            float32 `parquet:"name=mv_meter_phase_c_voltage_magnitude_Kv_sel735mv, type=DOUBLE" json:"mv_meter_phase_c_voltage_magnitude_Kv_sel735mv"`
	MVMeterPhaseABVoltageMagnitude           float32 `parquet:"name=mv_meter_phase_ab_voltage_magnitude_Kv_sel735mv, type=DOUBLE" json:"mv_meter_phase_ab_voltage_magnitude_Kv_sel735mv"`
	MVMeterPhaseBCVoltageMagnitude           float32 `parquet:"name=mv_meter_phase_bc_voltage_magnitude_Kv_sel735mv, type=DOUBLE" json:"mv_meter_phase_bc_voltage_magnitude_Kv_sel735mv"`
	MVMeterPhaseCAVoltageMagnitude           float32 `parquet:"name=mv_meter_phase_ca_voltage_magnitude_Kv_sel735mv, type=DOUBLE" json:"mv_meter_phase_ca_voltage_magnitude_Kv_sel735mv"`
	MVMeterPhaseACurrentAngle                float32 `parquet:"name=mv_meter_phase_a_current_angle_Deg_sel735mv, type=DOUBLE" json:"mv_meter_phase_a_current_angle_Deg_sel735mv"`
	MVMeterPhaseBCurrentAngle                float32 `parquet:"name=mv_meter_phase_b_current_angle_Deg_sel735mv, type=DOUBLE" json:"mv_meter_phase_b_current_angle_Deg_sel735mv"`
	MVMeterPhaseCCurrentAngle                float32 `parquet:"name=mv_meter_phase_c_current_angle_Deg_sel735mv, type=DOUBLE" json:"mv_meter_phase_c_current_angle_Deg_sel735mv"`
	MVMeterNeutralCurrentAngle               float32 `parquet:"name=mv_meter_neutral_current_angle_Deg_sel735mv, type=DOUBLE" json:"mv_meter_neutral_current_angle_Deg_sel735mv"`
	MVMeterPhaseAVoltageAngle                float32 `parquet:"name=mv_meter_phase_a_voltage_angle_Deg_sel735mv, type=DOUBLE" json:"mv_meter_phase_a_voltage_angle_Deg_sel735mv"`
	MVMeterPhaseBVoltageAngle                float32 `parquet:"name=mv_meter_phase_b_voltage_angle_Deg_sel735mv, type=DOUBLE" json:"mv_meter_phase_b_voltage_angle_Deg_sel735mv"`
	MVMeterPhaseCVoltageAngle                float32 `parquet:"name=mv_meter_phase_c_voltage_angle_Deg_sel735mv, type=DOUBLE" json:"mv_meter_phase_c_voltage_angle_Deg_sel735mv"`
	MVMeterPhaseABVoltageAngle               float32 `parquet:"name=mv_meter_phase_ab_voltage_angle_Deg_sel735mv, type=DOUBLE" json:"mv_meter_phase_ab_voltage_angle_Deg_sel735mv"`
	MVMeterPhaseBCVoltageAngle               float32 `parquet:"name=mv_meter_phase_bc_voltage_angle_Deg_sel735mv, type=DOUBLE" json:"mv_meter_phase_bc_voltage_angle_Deg_sel735mv"`
	MVMeterPhaseCAVoltageAngle               float32 `parquet:"name=mv_meter_phase_ca_voltage_angle_Deg_sel735mv, type=DOUBLE" json:"mv_meter_phase_ca_voltage_angle_Deg_sel735mv"`
	MVMeterDisplacementPowerFactor           float32 `parquet:"name=mv_meter_displacement_power_factor_sel735mv, type=DOUBLE" json:"mv_meter_displacement_power_factor_sel735mv"`
	HVMeterRealPowerDelivered                float32 `parquet:"name=hv_meter_real_power_delivered_KWh_sel735mv, type=DOUBLE" json:"hv_meter_real_power_delivered_KWh_sel735mv"`
	HVMeterRealPowerReceived                 float32 `parquet:"name=hv_meter_real_power_received_KWh_sel735mv, type=DOUBLE" json:"hv_meter_real_power_received_KWh_sel735mv"`
	HVMeterApparentPowerDelivered            float32 `parquet:"name=hv_meter_apparent_power_delivered_KVAh_sel735mv, type=DOUBLE" json:"hv_meter_apparent_power_delivered_KVAh_sel735mv"`
	HVMeterApparentPowerReceived             float32 `parquet:"name=hv_meter_apparent_power_received_KVAh_sel735mv, type=DOUBLE" json:"hv_meter_apparent_power_received_KVAh_sel735mv"`
	HVMeterReactivePowerDelivered            float32 `parquet:"name=hv_meter_reactive_power_delivered_KVarh_sel735mv, type=DOUBLE" json:"hv_meter_reactive_power_delivered_KVarh_sel735mv"`
	HVMeterReactivePowerReceived             float32 `parquet:"name=hv_meter_reactive_power_received_KVarh_sel735mv, type=DOUBLE" json:"hv_meter_reactive_power_received_KVarh_sel735mv"`
	HVMeterNetRealPowerDeliveredReceived     float32 `parquet:"name=hv_meter_net_real_power_delivered_received_KWh_sel735mv, type=DOUBLE" json:"hv_meter_net_real_power_delivered_received_KWh_sel735mv"`
	HVMeterNetReactivePowerDeliveredReceived float32 `parquet:"name=hv_meter_net_reactive_power_delivered_received_KVarh_sel735mv, type=DOUBLE" json:"hv_meter_net_reactive_power_delivered_received_KVarh_sel735mv"`
	ATSXfmrPhaseANVoltage                    float32 `parquet:"name=ats_xfmr_phase_an_voltage_V_abbtrueoneats, type=DOUBLE" json:"ats_xfmr_phase_an_voltage_V_abbtrueoneats"`
	ATSXfmrPhaseBNVoltage                    float32 `parquet:"name=ats_xfmr_phase_bn_voltage_V_abbtrueoneats, type=DOUBLE" json:"ats_xfmr_phase_bn_voltage_V_abbtrueoneats"`
	ATSXfmrPhaseABVoltage                    float32 `parquet:"name=ats_xfmr_phase_ab_voltage_V_abbtrueoneats, type=DOUBLE" json:"ats_xfmr_phase_ab_voltage_V_abbtrueoneats"`
	ATSGeneratorPhaseANVoltage               float32 `parquet:"name=ats_generator_phase_an_voltage_V_abbtrueoneats, type=DOUBLE" json:"ats_generator_phase_an_voltage_V_abbtrueoneats"`
	ATSGeneratorPhaseBNVoltage               float32 `parquet:"name=ats_generator_phase_bn_voltage_V_abbtrueoneats, type=DOUBLE" json:"ats_generator_phase_bn_voltage_V_abbtrueoneats"`
	ATSGeneratorPhaseABVoltage               float32 `parquet:"name=ats_generator_phase_ab_voltage_V_abbtrueoneats, type=DOUBLE" json:"ats_generator_phase_ab_voltage_V_abbtrueoneats"`
	ATSXfmrFrequency                         float32 `parquet:"name=ats_xfmr_frequency_HZ_abbtrueoneats, type=DOUBLE" json:"ats_xfmr_frequency_HZ_abbtrueoneats"`
	ATSGeneratorFrequency                    float32 `parquet:"name=ats_generator_frequency_HZ_abbtrueoneats, type=DOUBLE" json:"ats_generator_frequency_HZ_abbtrueoneats"`
	NumberOfTransfersForGensetToSupply       float32 `parquet:"name=number_of_transfers_for_genset_to_supply, type=DOUBLE" json:"number_of_transfers_for_genset_to_supply"`
	GeneratorFrequency                       float32 `parquet:"name=generator_frequency_HZ_genset, type=DOUBLE" json:"generator_frequency_HZ_genset"`
	GeneratorVoltageL3L1                     float32 `parquet:"name=generator_voltage_l3_l1_V_genset, type=DOUBLE" json:"generator_voltage_l3_l1_V_genset"`
	GeneratorCurrentPhaseL1                  float32 `parquet:"name=generator_current_phase_l1_A_genset, type=DOUBLE" json:"generator_current_phase_l1_A_genset"`
	GeneratorCurrentPhaseL3                  float32 `parquet:"name=generator_current_phase_l3_A_genset, type=DOUBLE" json:"generator_current_phase_l3_A_genset"`
	GeneratorVoltagePhaseL1N                 float32 `parquet:"name=generator_voltage_phase_l1_n_V_genset, type=DOUBLE" json:"generator_voltage_phase_l1_n_V_genset"`
	GeneratorVoltagePhaseL3N                 float32 `parquet:"name=generator_voltage_phase_l3_n_V_genset, type=DOUBLE" json:"generator_voltage_phase_l3_n_V_genset"`
	GeneratorFuelLevel                       float32 `parquet:"name=generator_fuel_level_genset, type=DOUBLE" json:"generator_fuel_level_genset"`
	GeneratorCoolantTemp                     float32 `parquet:"name=generator_coolant_temp_Degrees_C_genset, type=DOUBLE" json:"generator_coolant_temp_Degrees_C_genset"`
	GeneratorBatteryVoltage                  float32 `parquet:"name=generator_battery_voltage_V_genset, type=DOUBLE" json:"generator_battery_voltage_V_genset"`
	GeneratorEngineSpeed                     float32 `parquet:"name=generator_engine_speed_RPM_genset, type=DOUBLE" json:"generator_engine_speed_RPM_genset"`
	GeneratorEngineRunTime                   float32 `parquet:"name=generator_engine_run_time_SEC_genset, type=DOUBLE" json:"generator_engine_run_time_SEC_genset"`
	GeneratorStatus                          float32 `parquet:"name=generator_status_genset, type=DOUBLE" json:"generator_status_genset"`
	GeneratorSuccessStarts                   float32 `parquet:"name=generator_success_starts_genset, type=DOUBLE" json:"generator_success_starts_genset"`
	GeneratorOilPressure                     float32 `parquet:"name=generator_oil_pressure_KPA_genset, type=DOUBLE" json:"generator_oil_pressure_KPA_genset"`

	// source annotation to replace is `parquet:"name=@1@@, type=BOOLEAN" json:"@1@@"`
	H1ClosedStatus                         bool `parquet:"name=h1_closed_status_closed_ge845, type=BOOLEAN" json:"h1_closed_status_closed_ge845"`
	H1OpenStatus                           bool `parquet:"name=h1_open_status_open_ge845, type=BOOLEAN" json:"h1_open_status_open_ge845"`
	ClosedStatus89                         bool `parquet:"name=closed_status_closed_ge845, type=BOOLEAN" json:"closed_status_closed_ge845"`
	OpenStatus89                           bool `parquet:"name=open_status_open_ge845, type=BOOLEAN" json:"open_status_open_ge845"`
	A86T1A86BF52H1LockoutRelayOperated     bool `parquet:"name=a86t1___a86bf_52_h1_lockout_relay_operated_alarm_ge845, type=BOOLEAN" json:"a86t1___a86bf_52_h1_lockout_relay_operated_alarm_ge845"`
	XFMRT1TopOilTempAlarm                  bool `parquet:"name=xfmr_t1_top_oil_temp_alarm_ge845, type=BOOLEAN" json:"xfmr_t1_top_oil_temp_alarm_ge845"`
	XFMRT1WindingTempAlarm                 bool `parquet:"name=xfmr_t1_winding_temp_alarm_ge845, type=BOOLEAN" json:"xfmr_t1_winding_temp_alarm_ge845"`
	XFMRT1HighPressureAlarm                bool `parquet:"name=xfmr_t1_high_pressure_alarm_ge845, type=BOOLEAN" json:"xfmr_t1_high_pressure_alarm_ge845"`
	XFMRT1HighOilLevelMainAlarm            bool `parquet:"name=xfmr_t1_high_oil_level_main_alarm_ge845, type=BOOLEAN" json:"xfmr_t1_high_oil_level_main_alarm_ge845"`
	Trip52H1Coil1                          bool `parquet:"name=trip_52h1_coil_1_trip_ge845, type=BOOLEAN" json:"trip_52h1_coil_1_trip_ge845"`
	Trip52M1Coil1                          bool `parquet:"name=trip_52m1_coil_1_trip_ge845, type=BOOLEAN" json:"trip_52m1_coil_1_trip_ge845"`
	RelayFailureAlarmToSEL2240             bool `parquet:"name=relay_failure_alarm_to_sel2240_normal_ge845, type=BOOLEAN" json:"relay_failure_alarm_to_sel2240_normal_ge845"`
	TripLORA86T1                           bool `parquet:"name=trip_lor_a86t1_trip_ge845, type=BOOLEAN" json:"trip_lor_a86t1_trip_ge845"`
	MechanicalProtectionAlarm              bool `parquet:"name=mechanical_protection_alarm_ge845, type=BOOLEAN" json:"mechanical_protection_alarm_ge845"`
	Differential                           bool `parquet:"name=differential_alarm_ge845, type=BOOLEAN" json:"differential_alarm_ge845"`
	PhaseTimeOverCurrent                   bool `parquet:"name=phase_time_overcurrent_alarm_ge845, type=BOOLEAN" json:"phase_time_overcurrent_alarm_ge845"`
	GroundTimeOverCurrent                  bool `parquet:"name=ground_time_overcurrent_alarm_ge845, type=BOOLEAN" json:"ground_time_overcurrent_alarm_ge845"`
	VoltsPerHz                             bool `parquet:"name=volts_per_hz_alarm_ge845, type=BOOLEAN" json:"volts_per_hz_alarm_ge845"`
	PhaseUnderVoltage                      bool `parquet:"name=phase_undervoltage_alarm_ge845, type=BOOLEAN" json:"phase_undervoltage_alarm_ge845"`
	PhaseOverVoltage                       bool `parquet:"name=phase_overvoltage_alarm_ge845, type=BOOLEAN" json:"phase_overvoltage_alarm_ge845"`
	NeutralOverVoltage                     bool `parquet:"name=neutral_overvoltage_alarm_ge845, type=BOOLEAN" json:"neutral_overvoltage_alarm_ge845"`
	TimeSynch                              bool `parquet:"name=time_synch_alarm_ge845, type=BOOLEAN" json:"time_synch_alarm_ge845"`
	FuseFail                               bool `parquet:"name=fuse_fail_alarm_ge845, type=BOOLEAN" json:"fuse_fail_alarm_ge845"`
	CircuitBreakerNOContact                bool `parquet:"name=circuit_breaker_n_o_contact_closed_sel487e, type=BOOLEAN" json:"circuit_breaker_n_o_contact_closed_sel487e"`
	APhaseInvolvedInTheFault               bool `parquet:"name=a_phase_involved_in_the_fault_trip_sel487e, type=BOOLEAN" json:"a_phase_involved_in_the_fault_trip_sel487e"`
	BPhaseInvolvedInTheFault               bool `parquet:"name=b_phase_involved_in_the_fault_trip_sel487e, type=BOOLEAN" json:"b_phase_involved_in_the_fault_trip_sel487e"`
	CPhaseInvolvedInTheFault               bool `parquet:"name=c_phase_involved_in_the_fault_trip_sel487e, type=BOOLEAN" json:"c_phase_involved_in_the_fault_trip_sel487e"`
	BreakerTrip                            bool `parquet:"name=breaker_trip_trip_sel487e, type=BOOLEAN" json:"breaker_trip_trip_sel487e"`
	LossOfPotential                        bool `parquet:"name=loss_of_potential_alarm_sel487e, type=BOOLEAN" json:"loss_of_potential_alarm_sel487e"`
	TimeSourceAccuracy                     bool `parquet:"name=time_source_accuracy_normal_sel487e, type=BOOLEAN" json:"time_source_accuracy_normal_sel487e"`
	SoftwareAlarms                         bool `parquet:"name=software_alarms_alarm_sel487e, type=BOOLEAN" json:"software_alarms_alarm_sel487e"`
	HardwareAlarm                          bool `parquet:"name=hardware_alarm_sel487e, type=BOOLEAN" json:"hardware_alarm_sel487e"`
	BreakerH1Close                         bool `parquet:"name=breaker_h1_close_closed_sel487e, type=BOOLEAN" json:"breaker_h1_close_closed_sel487e"`
	BreakerH1Open                          bool `parquet:"name=breaker_h1_open_open_sel487e, type=BOOLEAN" json:"breaker_h1_open_open_sel487e"`
	L1CloseStatus89                        bool `parquet:"name=l1_close_status_closed_sel487e, type=BOOLEAN" json:"l1_close_status_closed_sel487e"`
	L1OpenStatus89                         bool `parquet:"name=l1_open_status_open_sel487e, type=BOOLEAN" json:"l1_open_status_open_sel487e"`
	LOROperate                             bool `parquet:"name=lor_operate_alarm_sel487e, type=BOOLEAN" json:"lor_operate_alarm_sel487e"`
	XFMRWindingTemperature                 bool `parquet:"name=xfmr_winding_temperature_alarm_sel487e, type=BOOLEAN" json:"xfmr_winding_temperature_alarm_sel487e"`
	XFMRLowPressure                        bool `parquet:"name=xfmr_low_pressure_alarm_sel487e, type=BOOLEAN" json:"xfmr_low_pressure_alarm_sel487e"`
	XFMRLowOilLevel                        bool `parquet:"name=xfmr_low_oil_level_alarm_sel487e, type=BOOLEAN" json:"xfmr_low_oil_level_alarm_sel487e"`
	TripBreakerH1                          bool `parquet:"name=trip_breaker_h1_trip_sel487e, type=BOOLEAN" json:"trip_breaker_h1_trip_sel487e"`
	TripBreakerM1                          bool `parquet:"name=trip_breaker_m1_trip_sel487e, type=BOOLEAN" json:"trip_breaker_m1_trip_sel487e"`
	TripB86T1                              bool `parquet:"name=trip_b86_t1_trip_sel487e, type=BOOLEAN" json:"trip_b86_t1_trip_sel487e"`
	TripB86BFH1                            bool `parquet:"name=trip_b86bf_h1_trip_sel487e, type=BOOLEAN" json:"trip_b86bf_h1_trip_sel487e"`
	RelayFailAlarm                         bool `parquet:"name=relay_fail_alarm_sel487e, type=BOOLEAN" json:"relay_fail_alarm_sel487e"`
	UnderFrequencyProtection               bool `parquet:"name=under_frequency_protection_trip_sel487e, type=BOOLEAN" json:"under_frequency_protection_trip_sel487e"`
	OverFrequencyProtection                bool `parquet:"name=over_frequency_protection_trip_sel487e, type=BOOLEAN" json:"over_frequency_protection_trip_sel487e"`
	DifferentialProtection                 bool `parquet:"name=differential_protection_trip_sel487e, type=BOOLEAN" json:"differential_protection_trip_sel487e"`
	HVWindingTimeOCProtection              bool `parquet:"name=hv_winding_time_oc_protection_trip_sel487e, type=BOOLEAN" json:"hv_winding_time_oc_protection_trip_sel487e"`
	MVWindingTimeOCProtection              bool `parquet:"name=mv_winding_time_oc_protection_trip_sel487e, type=BOOLEAN" json:"mv_winding_time_oc_protection_trip_sel487e"`
	VoltsPerHz2                            bool `parquet:"name=volts_per_hz_trip_sel487e, type=BOOLEAN" json:"volts_per_hz_trip_sel487e"`
	UnderVoltageProtection                 bool `parquet:"name=undervoltage_protection_trip_sel487e, type=BOOLEAN" json:"undervoltage_protection_trip_sel487e"`
	OverVoltageProtection                  bool `parquet:"name=overvoltage_protection_trip_sel487e, type=BOOLEAN" json:"overvoltage_protection_trip_sel487e"`
	OverVoltageNeutralProtection           bool `parquet:"name=overvoltage_neutral_protection_trip_sel487e, type=BOOLEAN" json:"overvoltage_neutral_protection_trip_sel487e"`
	HVMeterTimeBasedOnIRIGBTimeSource      bool `parquet:"name=hv_meter_time_based_on_irigb_time_source_normal_sel735hv, type=BOOLEAN" json:"hv_meter_time_based_on_irigb_time_source_normal_sel735hv"`
	HVMeterIRIGBInHighAccuracyMode         bool `parquet:"name=hv_meter_irigb_in_high_accuracy_mode_normal_sel735hv, type=BOOLEAN" json:"hv_meter_irigb_in_high_accuracy_mode_normal_sel735hv"`
	HVMeterSoftwareAlarm                   bool `parquet:"name=hv_meter_software_alarm_sel735hv, type=BOOLEAN" json:"hv_meter_software_alarm_sel735hv"`
	HVMeterHardwareAlarm                   bool `parquet:"name=hv_meter_hardware_alarm_sel735hv, type=BOOLEAN" json:"hv_meter_hardware_alarm_sel735hv"`
	HVMeterLDPFD3LeadLag                   bool `parquet:"name=hv_meter_ldpfd3_leadlag_lead_sel735hv, type=BOOLEAN" json:"hv_meter_ldpfd3_leadlag_lead_sel735hv"`
	HVMeterMainBoardOutput3                bool `parquet:"name=hv_meter_main_board_output_3_normal_sel735hv, type=BOOLEAN" json:"hv_meter_main_board_output_3_normal_sel735hv"`
	MVMeterTimeBasedOnIRIGBTimeSource      bool `parquet:"name=mv_meter_time_based_on_irigb_time_source_normal_sel735mv, type=BOOLEAN" json:"mv_meter_time_based_on_irigb_time_source_normal_sel735mv"`
	MVMeterIRIGBInHighAccuracyMode         bool `parquet:"name=mv_meter_irigb_in_high_accuracy_mode_normal_sel735mv, type=BOOLEAN" json:"mv_meter_irigb_in_high_accuracy_mode_normal_sel735mv"`
	MVMeterSoftwareAlarm                   bool `parquet:"name=mv_meter_software_alarm_sel735mv, type=BOOLEAN" json:"mv_meter_software_alarm_sel735mv"`
	MVMeterHardwareAlarm                   bool `parquet:"name=mv_meter_hardware_alarm_sel735mv, type=BOOLEAN" json:"mv_meter_hardware_alarm_sel735mv"`
	MVMeterLDPFD3LeadLag                   bool `parquet:"name=mv_meter_ldpfd3_leadlag_lead_sel735mv, type=BOOLEAN" json:"mv_meter_ldpfd3_leadlag_lead_sel735mv"`
	MVMeterMainBoardOutput3                bool `parquet:"name=mv_meter_main_board_output_3_normal_sel735mv, type=BOOLEAN" json:"mv_meter_main_board_output_3_normal_sel735mv"`
	CommonFailAlarm                        bool `parquet:"name=common_fail_alarm_abbats, type=BOOLEAN" json:"common_fail_alarm_abbats"`
	ATSInAutoMode                          bool `parquet:"name=ats_in_auto_mode_normal_abbats, type=BOOLEAN" json:"ats_in_auto_mode_normal_abbats"`
	ATSConnectedToTransformer              bool `parquet:"name=ats_connected_to_transformer_normal_abbats, type=BOOLEAN" json:"ats_connected_to_transformer_normal_abbats"`
	A86BF52T1LORTroubleAlarm               bool `parquet:"name=a86bf52t1_lor_trouble_alarm_axion, type=BOOLEAN" json:"a86bf52t1_lor_trouble_alarm_axion"`
	A86T1LORTroubleAlarm                   bool `parquet:"name=a86t1_lor_trouble_alarm_axion, type=BOOLEAN" json:"a86t1_lor_trouble_alarm_axion"`
	B86BF52T1LORTroubleAlarm               bool `parquet:"name=b86bf52t1_lor_trouble_alarm_axion, type=BOOLEAN" json:"b86bf52t1_lor_trouble_alarm_axion"`
	B86T1LORTroubleAlarm                   bool `parquet:"name=b86t1_lor_trouble_alarm_axion, type=BOOLEAN" json:"b86t1_lor_trouble_alarm_axion"`
	A86BF52T1LOROperated                   bool `parquet:"name=a86bf52t1_lor_operated_alarm_axion, type=BOOLEAN" json:"a86bf52t1_lor_operated_alarm_axion"`
	A86T1LOROperated                       bool `parquet:"name=a86t1_lor_operated_alarm_axion, type=BOOLEAN" json:"a86t1_lor_operated_alarm_axion"`
	B86BF52T1LOROperated                   bool `parquet:"name=b86bf52t1_lor_operated_alarm_axion, type=BOOLEAN" json:"b86bf52t1_lor_operated_alarm_axion"`
	B86T1LOROperated                       bool `parquet:"name=b86t1_lor_operated_alarm_axion, type=BOOLEAN" json:"b86t1_lor_operated_alarm_axion"`
	GE845Normal                            bool `parquet:"name=ge_845_normal_normal_axion, type=BOOLEAN" json:"ge_845_normal_normal_axion"`
	SEL487ENormal                          bool `parquet:"name=sel487e_normal_normal_axion, type=BOOLEAN" json:"sel487e_normal_normal_axion"`
	SEL2730MNormal                         bool `parquet:"name=sel2730m_normal_normal_axion, type=BOOLEAN" json:"sel2730m_normal_normal_axion"`
	SEL2401Normal                          bool `parquet:"name=sel2401_normal_normal_axion, type=BOOLEAN" json:"sel2401_normal_normal_axion"`
	SEL735HVMETERNormal                    bool `parquet:"name=sel735_hvmeter_normal_normal_axion, type=BOOLEAN" json:"sel735_hvmeter_normal_normal_axion"`
	SEL735MVMETERNormal                    bool `parquet:"name=sel735_mvmeter_normal_normal_axion, type=BOOLEAN" json:"sel735_mvmeter_normal_normal_axion"`
	InverterOutputNormal                   bool `parquet:"name=inverter_output_normal_normal_axion, type=BOOLEAN" json:"inverter_output_normal_normal_axion"`
	HVACFailAlarm                          bool `parquet:"name=hvac_fail_alarm_axion, type=BOOLEAN" json:"hvac_fail_alarm_axion"`
	Section1BattChgrAlarm                  bool `parquet:"name=section_1_batt_chgr_alarm_normal_axion, type=BOOLEAN" json:"section_1_batt_chgr_alarm_normal_axion"`
	Section3BattChgrAlarm                  bool `parquet:"name=section_3_batt_chgr_alarm_normal_axion, type=BOOLEAN" json:"section_3_batt_chgr_alarm_normal_axion"`
	L1Closed89                             bool `parquet:"name=l1_closed_open_axion, type=BOOLEAN" json:"l1_closed_open_axion"`
	L1Open89                               bool `parquet:"name=l1_open_closed_axion, type=BOOLEAN" json:"l1_open_closed_axion"`
	L1LossOfACAlarm89                      bool `parquet:"name=l1_loss_of_ac_alarm_axion, type=BOOLEAN" json:"l1_loss_of_ac_alarm_axion"`
	L1LossOfDCAlarm89                      bool `parquet:"name=l1_loss_of_dc_alarm_axion, type=BOOLEAN" json:"l1_loss_of_dc_alarm_axion"`
	H1Closed52                             bool `parquet:"name=h1_closed_open_axion, type=BOOLEAN" json:"h1_closed_open_axion"`
	H1Open52                               bool `parquet:"name=h1_open_closed_axion, type=BOOLEAN" json:"h1_open_closed_axion"`
	H1LowSF6PressAlarm52                   bool `parquet:"name=h1_low_sf6_press_alarm_normal_axion, type=BOOLEAN" json:"h1_low_sf6_press_alarm_normal_axion"`
	H1LowSF6PressCutout52                  bool `parquet:"name=h1_low_sf6_press_cutout_normal_axion, type=BOOLEAN" json:"h1_low_sf6_press_cutout_normal_axion"`
	XFMRT1TopOilTempAlarm70C               bool `parquet:"name=xfmr_t1_top_oil_temp_alarm_70c_alarm_axion, type=BOOLEAN" json:"xfmr_t1_top_oil_temp_alarm_70c_alarm_axion"`
	XFMRT1WindingTempAlarm75C              bool `parquet:"name=xfmr_t1_winding_temp_alarm_75c_alarm_axion, type=BOOLEAN" json:"xfmr_t1_winding_temp_alarm_75c_alarm_axion"`
	XFMRT1LowN2CylinderPressAlarm200PSI    bool `parquet:"name=xfmr_t1_low_n2_cylinder_press_alarm_200_psi_alarm_axion, type=BOOLEAN" json:"xfmr_t1_low_n2_cylinder_press_alarm_200_psi_alarm_axion"`
	XFMRT1PressReliefAlarmMainTank         bool `parquet:"name=xfmr_t1_press_relief_alarm_main_tank_alarm_axion, type=BOOLEAN" json:"xfmr_t1_press_relief_alarm_main_tank_alarm_axion"`
	XFMRT1LossOfVoltageMainSupplyAlarm     bool `parquet:"name=xfmr_t1_loss_of_voltage_main_supply_alarm_axion, type=BOOLEAN" json:"xfmr_t1_loss_of_voltage_main_supply_alarm_axion"`
	XFMRT1LossOfVoltageFanSupplyAlarm      bool `parquet:"name=xfmr_t1_loss_of_voltage_fan_supply_alarm_axion, type=BOOLEAN" json:"xfmr_t1_loss_of_voltage_fan_supply_alarm_axion"`
	XFMRT1LossOfVoltageCoolingControlAlarm bool `parquet:"name=xfmr_t1_loss_of_voltage_cooling_control_alarm_axion, type=BOOLEAN" json:"xfmr_t1_loss_of_voltage_cooling_control_alarm_axion"`
	T1Closed89                             bool `parquet:"name=t1_closed_open_axion, type=BOOLEAN" json:"t1_closed_open_axion"`
	T1Open89                               bool `parquet:"name=t1_open_closed_axion, type=BOOLEAN" json:"t1_open_closed_axion"`
	T1LossOfACAlarm89                      bool `parquet:"name=t1_loss_of_ac_alarm_axion, type=BOOLEAN" json:"t1_loss_of_ac_alarm_axion"`
	T1LossOfDCAlarm89                      bool `parquet:"name=t1_loss_of_dc_alarm_axion, type=BOOLEAN" json:"t1_loss_of_dc_alarm_axion"`
}

func ModbusRegistersFomRegisters(ri *InputRegister, rd *DiscreteRegister) ModbusRegisters {
	tn := time.Now()
	return ModbusRegisters{
		TimestampPQ:                              tn.Format(time.DateTime),
		TimestampNum:                             utils.TimeToDayIntNum(tn),
		InstantaneousTemp:                        ri.InstantaneousTemp,
		AverageHumidity:                          ri.AverageHumidity,
		MaximumHumidity:                          ri.MaximumHumidity,
		MinimumHumidity:                          ri.MinimumHumidity,
		AverageAmbientTemperature:                ri.AverageAmbientTemperature,
		MaximumAmbientTemperature:                ri.MaximumAmbientTemperature,
		MinimumAmbientTemperature:                ri.MinimumAmbientTemperature,
		PhaseACurrent145kV:                       ri.PhaseACurrent145kV,
		PhaseBCurrent145kV:                       ri.PhaseBCurrent145kV,
		PhaseCCurrent145kV:                       ri.PhaseCCurrent145kV,
		PhaseCCurrent38kV:                        ri.PhaseCCurrent38kV,
		PhaseACurrent38kV:                        ri.PhaseACurrent38kV,
		PhaseBCurrent38kV:                        ri.PhaseBCurrent38kV,
		PhaseABVoltage145kV:                      ri.PhaseABVoltage145kV,
		PhaseBCVoltage145kV:                      ri.PhaseBCVoltage145kV,
		PhaseCAVoltage145kV:                      ri.PhaseCAVoltage145kV,
		RealPower:                                ri.RealPower,
		ReactivePower:                            ri.ReactivePower,
		ApparentPower:                            ri.ApparentPower,
		PowerFactor:                              ri.PowerFactor,
		Frequency:                                ri.Frequency,
		SidePhaseAN145kV:                         ri.SidePhaseAN145kV,
		SidePhaseBN145kV:                         ri.SidePhaseBN145kV,
		SidePhaseCN145kV:                         ri.SidePhaseCN145kV,
		SideNeutralCurrent38kV:                   ri.SideNeutralCurrent38kV,
		SidePhaseAVoltage145kV:                   ri.SidePhaseAVoltage145kV,
		SidePhaseBVoltage145kV:                   ri.SidePhaseBVoltage145kV,
		SidePhaseCVoltage145kV:                   ri.SidePhaseCVoltage145kV,
		SidePhaseABVoltage145kV:                  ri.SidePhaseABVoltage145kV,
		SidePhaseBCVoltage145kV:                  ri.SidePhaseBCVoltage145kV,
		SidePhaseCAVoltage145kV:                  ri.SidePhaseCAVoltage145kV,
		SidePhaseACurrent145kV:                   ri.SidePhaseACurrent145kV,
		SidePhaseBCurrent145kV:                   ri.SidePhaseBCurrent145kV,
		SidePhaseCCurrent145kV:                   ri.SidePhaseCCurrent145kV,
		SidePhaseACurrent38kV:                    ri.SidePhaseACurrent38kV,
		SidePhaseBCurrent38kV:                    ri.SidePhaseBCurrent38kV,
		SidePhaseCCurrent38kV:                    ri.SidePhaseCCurrent38kV,
		SideRealPower:                            ri.SideRealPower,
		SideReactivePower:                        ri.SideReactivePower,
		SideApparentPower:                        ri.SideApparentPower,
		Frequency2:                               ri.Frequency2,
		PhasePowerFactorDisplacement:             ri.PhasePowerFactorDisplacement,
		PhaseARMSCurrent:                         ri.PhaseARMSCurrent,
		PhaseBRMSCurrent:                         ri.PhaseBRMSCurrent,
		PhaseCRMSCurrent:                         ri.PhaseCRMSCurrent,
		NeutralRMSCurrent:                        ri.NeutralRMSCurrent,
		PhaseARMSVoltage:                         ri.PhaseARMSVoltage,
		PhaseBRMSVoltage:                         ri.PhaseBRMSVoltage,
		PhaseCRMSVoltage:                         ri.PhaseCRMSVoltage,
		PhaseABRMSVoltage:                        ri.PhaseABRMSVoltage,
		PhaseBCRMSVoltage:                        ri.PhaseBCRMSVoltage,
		PhaseCARMSVoltage:                        ri.PhaseCARMSVoltage,
		RealPower2:                               ri.RealPower2,
		ApparentPower2:                           ri.ApparentPower2,
		ReactivePower2:                           ri.ReactivePower2,
		PhaseARealPower:                          ri.PhaseARealPower,
		PhaseBRealPower:                          ri.PhaseBRealPower,
		PhaseCRealPower:                          ri.PhaseCRealPower,
		PhaseAApparentPower:                      ri.PhaseAApparentPower,
		PhaseBApparentPower:                      ri.PhaseBApparentPower,
		PhaseCApparentPower:                      ri.PhaseCApparentPower,
		PhaseAReactivePower:                      ri.PhaseAReactivePower,
		PhaseBReactivePower:                      ri.PhaseBReactivePower,
		PhaseCReactivePower:                      ri.PhaseCReactivePower,
		Frequency3:                               ri.Frequency3,
		PhaseACurrentMagnitude:                   ri.PhaseACurrentMagnitude,
		PhaseBCurrentMagnitude:                   ri.PhaseBCurrentMagnitude,
		PhaseCCurrentMagnitude:                   ri.PhaseCCurrentMagnitude,
		NeutralCurrentMagnitude:                  ri.NeutralCurrentMagnitude,
		PhaseAVoltageMagnitude:                   ri.PhaseAVoltageMagnitude,
		PhaseBVoltageMagnitude:                   ri.PhaseBVoltageMagnitude,
		PhaseCVoltageMagnitude:                   ri.PhaseCVoltageMagnitude,
		PhaseABVoltageMagnitude:                  ri.PhaseABVoltageMagnitude,
		PhaseBCVoltageMagnitude:                  ri.PhaseBCVoltageMagnitude,
		PhaseCAVoltageMagnitude:                  ri.PhaseCAVoltageMagnitude,
		PhaseACurrentAngle:                       ri.PhaseACurrentAngle,
		PhaseBCurrentAngle:                       ri.PhaseBCurrentAngle,
		PhaseCCurrentAngle:                       ri.PhaseCCurrentAngle,
		NeutralCurrentAngle:                      ri.NeutralCurrentAngle,
		PhaseAVoltageAngle:                       ri.PhaseAVoltageAngle,
		PhaseBVoltageAngle:                       ri.PhaseBVoltageAngle,
		PhaseCVoltageAngle:                       ri.PhaseCVoltageAngle,
		PhaseABVoltageAngle:                      ri.PhaseABVoltageAngle,
		PhaseBCVoltageAngle:                      ri.PhaseBCVoltageAngle,
		PhaseCAVoltageAngle:                      ri.PhaseCAVoltageAngle,
		DisplacementPowerFactor:                  ri.DisplacementPowerFactor,
		RealPowerDelivered:                       ri.RealPowerDelivered,
		RealPowerReceived:                        ri.RealPowerReceived,
		ApparentPowerDelivered:                   ri.ApparentPowerDelivered,
		ApparentPowerReceived:                    ri.ApparentPowerReceived,
		ReactivePowerDelivered:                   ri.ReactivePowerDelivered,
		ReactivePowerReceived:                    ri.ReactivePowerReceived,
		NetRealPowerDeliveredReceived:            ri.NetRealPowerDeliveredReceived,
		NetReactivePowerDeliveredReceived:        ri.NetReactivePowerDeliveredReceived,
		MVMeterPhaseARMSCurrent:                  ri.MVMeterPhaseARMSCurrent,
		MVMeterPhaseBRMSCurrent:                  ri.MVMeterPhaseBRMSCurrent,
		MVMeterPhaseCRMSCurrent:                  ri.MVMeterPhaseCRMSCurrent,
		MVMeterNeutralRMSCurrent:                 ri.MVMeterNeutralRMSCurrent,
		MVMeterPhaseARMSVoltage:                  ri.MVMeterPhaseARMSVoltage,
		MVMeterPhaseBRMSVoltage:                  ri.MVMeterPhaseBRMSVoltage,
		MVMeterPhaseCRMSVoltage:                  ri.MVMeterPhaseCRMSVoltage,
		MVMeterPhaseABRMSVoltage:                 ri.MVMeterPhaseABRMSVoltage,
		MVMeterPhaseBCRMSVoltage:                 ri.MVMeterPhaseBCRMSVoltage,
		MVMeterPhaseCARMSVoltage:                 ri.MVMeterPhaseCARMSVoltage,
		MVMeterRealPower:                         ri.MVMeterRealPower,
		MVMeterApparentPower:                     ri.MVMeterApparentPower,
		MVMeterReactivePower:                     ri.MVMeterReactivePower,
		MVMeterPhaseARealPower:                   ri.MVMeterPhaseARealPower,
		MVMeterPhaseBRealPower:                   ri.MVMeterPhaseBRealPower,
		MVMeterPhaseCRealPower:                   ri.MVMeterPhaseCRealPower,
		MVMeterPhaseAApparentPower:               ri.MVMeterPhaseAApparentPower,
		MVMeterPhaseBApparentPower:               ri.MVMeterPhaseBApparentPower,
		MVMeterPhaseCApparentPower:               ri.MVMeterPhaseCApparentPower,
		MVMeterPhaseAReactivePower:               ri.MVMeterPhaseAReactivePower,
		MVMeterPhaseBReactivePower:               ri.MVMeterPhaseBReactivePower,
		MVMeterPhaseCReactivePower:               ri.MVMeterPhaseCReactivePower,
		MVMeterFrequency:                         ri.MVMeterFrequency,
		MVMeterPhaseACurrentMagnitude:            ri.MVMeterPhaseACurrentMagnitude,
		MVMeterPhaseBCurrentMagnitude:            ri.MVMeterPhaseBCurrentMagnitude,
		MVMeterPhaseCCurrentMagnitude:            ri.MVMeterPhaseCCurrentMagnitude,
		MVMeterNeutralCurrentMagnitude:           ri.MVMeterNeutralCurrentMagnitude,
		MVMeterPhaseAVoltageMagnitude:            ri.MVMeterPhaseAVoltageMagnitude,
		MVMeterPhaseBVoltageMagnitude:            ri.MVMeterPhaseBVoltageMagnitude,
		MVMeterPhaseCVoltageMagnitude:            ri.MVMeterPhaseCVoltageMagnitude,
		MVMeterPhaseABVoltageMagnitude:           ri.MVMeterPhaseABVoltageMagnitude,
		MVMeterPhaseBCVoltageMagnitude:           ri.MVMeterPhaseBCVoltageMagnitude,
		MVMeterPhaseCAVoltageMagnitude:           ri.MVMeterPhaseCAVoltageMagnitude,
		MVMeterPhaseACurrentAngle:                ri.MVMeterPhaseACurrentAngle,
		MVMeterPhaseBCurrentAngle:                ri.MVMeterPhaseBCurrentAngle,
		MVMeterPhaseCCurrentAngle:                ri.MVMeterPhaseCCurrentAngle,
		MVMeterNeutralCurrentAngle:               ri.MVMeterNeutralCurrentAngle,
		MVMeterPhaseAVoltageAngle:                ri.MVMeterPhaseAVoltageAngle,
		MVMeterPhaseBVoltageAngle:                ri.MVMeterPhaseBVoltageAngle,
		MVMeterPhaseCVoltageAngle:                ri.MVMeterPhaseCVoltageAngle,
		MVMeterPhaseABVoltageAngle:               ri.MVMeterPhaseABVoltageAngle,
		MVMeterPhaseBCVoltageAngle:               ri.MVMeterPhaseBCVoltageAngle,
		MVMeterPhaseCAVoltageAngle:               ri.MVMeterPhaseCAVoltageAngle,
		MVMeterDisplacementPowerFactor:           ri.MVMeterDisplacementPowerFactor,
		HVMeterRealPowerDelivered:                ri.HVMeterRealPowerDelivered,
		HVMeterRealPowerReceived:                 ri.HVMeterRealPowerReceived,
		HVMeterApparentPowerDelivered:            ri.HVMeterApparentPowerDelivered,
		HVMeterApparentPowerReceived:             ri.HVMeterApparentPowerReceived,
		HVMeterReactivePowerDelivered:            ri.HVMeterReactivePowerDelivered,
		HVMeterReactivePowerReceived:             ri.HVMeterReactivePowerReceived,
		HVMeterNetRealPowerDeliveredReceived:     ri.HVMeterNetRealPowerDeliveredReceived,
		HVMeterNetReactivePowerDeliveredReceived: ri.HVMeterNetReactivePowerDeliveredReceived,
		ATSXfmrPhaseANVoltage:                    ri.ATSXfmrPhaseANVoltage,
		ATSXfmrPhaseBNVoltage:                    ri.ATSXfmrPhaseBNVoltage,
		ATSXfmrPhaseABVoltage:                    ri.ATSXfmrPhaseABVoltage,
		ATSGeneratorPhaseANVoltage:               ri.ATSGeneratorPhaseANVoltage,
		ATSGeneratorPhaseBNVoltage:               ri.ATSGeneratorPhaseBNVoltage,
		ATSGeneratorPhaseABVoltage:               ri.ATSGeneratorPhaseABVoltage,
		ATSXfmrFrequency:                         ri.ATSXfmrFrequency,
		ATSGeneratorFrequency:                    ri.ATSGeneratorFrequency,
		NumberOfTransfersForGensetToSupply:       ri.NumberOfTransfersForGensetToSupply,
		GeneratorFrequency:                       ri.GeneratorFrequency,
		GeneratorVoltageL3L1:                     ri.GeneratorVoltageL3L1,
		GeneratorCurrentPhaseL1:                  ri.GeneratorCurrentPhaseL1,
		GeneratorCurrentPhaseL3:                  ri.GeneratorCurrentPhaseL3,
		GeneratorVoltagePhaseL1N:                 ri.GeneratorVoltagePhaseL1N,
		GeneratorVoltagePhaseL3N:                 ri.GeneratorVoltagePhaseL3N,
		GeneratorFuelLevel:                       ri.GeneratorFuelLevel,
		GeneratorCoolantTemp:                     ri.GeneratorCoolantTemp,
		GeneratorBatteryVoltage:                  ri.GeneratorBatteryVoltage,
		GeneratorEngineSpeed:                     ri.GeneratorEngineSpeed,
		GeneratorEngineRunTime:                   ri.GeneratorEngineRunTime,
		GeneratorStatus:                          ri.GeneratorStatus,
		GeneratorSuccessStarts:                   ri.GeneratorSuccessStarts,
		GeneratorOilPressure:                     ri.GeneratorOilPressure,
		H1ClosedStatus:                           rd.H1ClosedStatus,
		H1OpenStatus:                             rd.H1OpenStatus,
		ClosedStatus89:                           rd.ClosedStatus89,
		OpenStatus89:                             rd.OpenStatus89,
		A86T1A86BF52H1LockoutRelayOperated:       rd.A86T1A86BF52H1LockoutRelayOperated,
		XFMRT1TopOilTempAlarm:                    rd.XFMRT1TopOilTempAlarm,
		XFMRT1WindingTempAlarm:                   rd.XFMRT1WindingTempAlarm,
		XFMRT1HighPressureAlarm:                  rd.XFMRT1HighPressureAlarm,
		XFMRT1HighOilLevelMainAlarm:              rd.XFMRT1HighOilLevelMainAlarm,
		Trip52H1Coil1:                            rd.Trip52H1Coil1,
		Trip52M1Coil1:                            rd.Trip52M1Coil1,
		RelayFailureAlarmToSEL2240:               rd.RelayFailureAlarmToSEL2240,
		TripLORA86T1:                             rd.TripLORA86T1,
		MechanicalProtectionAlarm:                rd.MechanicalProtectionAlarm,
		Differential:                             rd.Differential,
		PhaseTimeOverCurrent:                     rd.PhaseTimeOverCurrent,
		GroundTimeOverCurrent:                    rd.GroundTimeOverCurrent,
		VoltsPerHz:                               rd.VoltsPerHz,
		PhaseUnderVoltage:                        rd.PhaseUnderVoltage,
		PhaseOverVoltage:                         rd.PhaseOverVoltage,
		NeutralOverVoltage:                       rd.NeutralOverVoltage,
		TimeSynch:                                rd.TimeSynch,
		FuseFail:                                 rd.FuseFail,
		CircuitBreakerNOContact:                  rd.CircuitBreakerNOContact,
		APhaseInvolvedInTheFault:                 rd.APhaseInvolvedInTheFault,
		BPhaseInvolvedInTheFault:                 rd.BPhaseInvolvedInTheFault,
		CPhaseInvolvedInTheFault:                 rd.CPhaseInvolvedInTheFault,
		BreakerTrip:                              rd.BreakerTrip,
		LossOfPotential:                          rd.LossOfPotential,
		TimeSourceAccuracy:                       rd.TimeSourceAccuracy,
		SoftwareAlarms:                           rd.SoftwareAlarms,
		HardwareAlarm:                            rd.HardwareAlarm,
		BreakerH1Close:                           rd.BreakerH1Close,
		BreakerH1Open:                            rd.BreakerH1Open,
		L1CloseStatus89:                          rd.L1CloseStatus89,
		L1OpenStatus89:                           rd.L1OpenStatus89,
		LOROperate:                               rd.LOROperate,
		XFMRWindingTemperature:                   rd.XFMRWindingTemperature,
		XFMRLowPressure:                          rd.XFMRLowPressure,
		XFMRLowOilLevel:                          rd.XFMRLowOilLevel,
		TripBreakerH1:                            rd.TripBreakerH1,
		TripBreakerM1:                            rd.TripBreakerM1,
		TripB86T1:                                rd.TripB86T1,
		TripB86BFH1:                              rd.TripB86BFH1,
		RelayFailAlarm:                           rd.RelayFailAlarm,
		UnderFrequencyProtection:                 rd.UnderFrequencyProtection,
		OverFrequencyProtection:                  rd.OverFrequencyProtection,
		DifferentialProtection:                   rd.DifferentialProtection,
		HVWindingTimeOCProtection:                rd.HVWindingTimeOCProtection,
		MVWindingTimeOCProtection:                rd.MVWindingTimeOCProtection,
		VoltsPerHz2:                              rd.VoltsPerHz2,
		UnderVoltageProtection:                   rd.UnderVoltageProtection,
		OverVoltageProtection:                    rd.OverVoltageProtection,
		OverVoltageNeutralProtection:             rd.OverVoltageNeutralProtection,
		HVMeterTimeBasedOnIRIGBTimeSource:        rd.HVMeterTimeBasedOnIRIGBTimeSource,
		HVMeterIRIGBInHighAccuracyMode:           rd.HVMeterIRIGBInHighAccuracyMode,
		HVMeterSoftwareAlarm:                     rd.HVMeterSoftwareAlarm,
		HVMeterHardwareAlarm:                     rd.HVMeterHardwareAlarm,
		HVMeterLDPFD3LeadLag:                     rd.HVMeterLDPFD3LeadLag,
		HVMeterMainBoardOutput3:                  rd.HVMeterMainBoardOutput3,
		MVMeterTimeBasedOnIRIGBTimeSource:        rd.MVMeterTimeBasedOnIRIGBTimeSource,
		MVMeterIRIGBInHighAccuracyMode:           rd.MVMeterIRIGBInHighAccuracyMode,
		MVMeterSoftwareAlarm:                     rd.MVMeterSoftwareAlarm,
		MVMeterHardwareAlarm:                     rd.MVMeterHardwareAlarm,
		MVMeterLDPFD3LeadLag:                     rd.MVMeterLDPFD3LeadLag,
		MVMeterMainBoardOutput3:                  rd.MVMeterMainBoardOutput3,
		CommonFailAlarm:                          rd.CommonFailAlarm,
		ATSInAutoMode:                            rd.ATSInAutoMode,
		ATSConnectedToTransformer:                rd.ATSConnectedToTransformer,
		A86BF52T1LORTroubleAlarm:                 rd.A86BF52T1LORTroubleAlarm,
		A86T1LORTroubleAlarm:                     rd.A86T1LORTroubleAlarm,
		B86BF52T1LORTroubleAlarm:                 rd.B86BF52T1LORTroubleAlarm,
		B86T1LORTroubleAlarm:                     rd.B86T1LORTroubleAlarm,
		A86BF52T1LOROperated:                     rd.A86BF52T1LOROperated,
		A86T1LOROperated:                         rd.A86T1LOROperated,
		B86BF52T1LOROperated:                     rd.B86BF52T1LOROperated,
		B86T1LOROperated:                         rd.B86T1LOROperated,
		GE845Normal:                              rd.GE845Normal,
		SEL487ENormal:                            rd.SEL487ENormal,
		SEL2730MNormal:                           rd.SEL2730MNormal,
		SEL2401Normal:                            rd.SEL2401Normal,
		SEL735HVMETERNormal:                      rd.SEL735HVMETERNormal,
		SEL735MVMETERNormal:                      rd.SEL735MVMETERNormal,
		InverterOutputNormal:                     rd.InverterOutputNormal,
		HVACFailAlarm:                            rd.HVACFailAlarm,
		Section1BattChgrAlarm:                    rd.Section1BattChgrAlarm,
		Section3BattChgrAlarm:                    rd.Section3BattChgrAlarm,
		L1Closed89:                               rd.L1Closed89,
		L1Open89:                                 rd.L1Open89,
		L1LossOfACAlarm89:                        rd.L1LossOfACAlarm89,
		L1LossOfDCAlarm89:                        rd.L1LossOfDCAlarm89,
		H1Closed52:                               rd.H1Closed52,
		H1Open52:                                 rd.H1Open52,
		H1LowSF6PressAlarm52:                     rd.H1LowSF6PressAlarm52,
		H1LowSF6PressCutout52:                    rd.H1LowSF6PressCutout52,
		XFMRT1TopOilTempAlarm70C:                 rd.XFMRT1TopOilTempAlarm70C,
		XFMRT1WindingTempAlarm75C:                rd.XFMRT1WindingTempAlarm75C,
		XFMRT1LowN2CylinderPressAlarm200PSI:      rd.XFMRT1LowN2CylinderPressAlarm200PSI,
		XFMRT1PressReliefAlarmMainTank:           rd.XFMRT1PressReliefAlarmMainTank,
		XFMRT1LossOfVoltageMainSupplyAlarm:       rd.XFMRT1LossOfVoltageMainSupplyAlarm,
		XFMRT1LossOfVoltageFanSupplyAlarm:        rd.XFMRT1LossOfVoltageFanSupplyAlarm,
		XFMRT1LossOfVoltageCoolingControlAlarm:   rd.XFMRT1LossOfVoltageCoolingControlAlarm,
		T1Closed89:                               rd.T1Closed89,
		T1Open89:                                 rd.T1Open89,
		T1LossOfACAlarm89:                        rd.T1LossOfACAlarm89,
		T1LossOfDCAlarm89:                        rd.T1LossOfDCAlarm89,
	}
}

type InputRegister struct {
	InstantaneousTemp                        float32
	AverageHumidity                          float32
	MaximumHumidity                          float32
	MinimumHumidity                          float32
	AverageAmbientTemperature                float32
	MaximumAmbientTemperature                float32
	MinimumAmbientTemperature                float32
	PhaseACurrent145kV                       float32
	PhaseBCurrent145kV                       float32
	PhaseCCurrent145kV                       float32
	PhaseACurrent38kV                        float32
	PhaseBCurrent38kV                        float32
	PhaseCCurrent38kV                        float32
	PhaseABVoltage145kV                      float32
	PhaseBCVoltage145kV                      float32
	PhaseCAVoltage145kV                      float32
	RealPower                                float32
	ReactivePower                            float32
	ApparentPower                            float32
	PowerFactor                              float32
	Frequency                                float32
	SidePhaseAN145kV                         float32
	SidePhaseBN145kV                         float32
	SidePhaseCN145kV                         float32
	SideNeutralCurrent38kV                   float32
	SidePhaseAVoltage145kV                   float32
	SidePhaseBVoltage145kV                   float32
	SidePhaseCVoltage145kV                   float32
	SidePhaseABVoltage145kV                  float32
	SidePhaseBCVoltage145kV                  float32
	SidePhaseCAVoltage145kV                  float32
	SidePhaseACurrent145kV                   float32
	SidePhaseBCurrent145kV                   float32
	SidePhaseCCurrent145kV                   float32
	SidePhaseACurrent38kV                    float32
	SidePhaseBCurrent38kV                    float32
	SidePhaseCCurrent38kV                    float32
	SideRealPower                            float32
	SideReactivePower                        float32
	SideApparentPower                        float32
	Frequency2                               float32
	PhasePowerFactorDisplacement             float32
	PhaseARMSCurrent                         float32
	PhaseBRMSCurrent                         float32
	PhaseCRMSCurrent                         float32
	NeutralRMSCurrent                        float32
	PhaseARMSVoltage                         float32
	PhaseBRMSVoltage                         float32
	PhaseCRMSVoltage                         float32
	PhaseABRMSVoltage                        float32
	PhaseBCRMSVoltage                        float32
	PhaseCARMSVoltage                        float32
	RealPower2                               float32
	ApparentPower2                           float32
	ReactivePower2                           float32
	PhaseARealPower                          float32
	PhaseBRealPower                          float32
	PhaseCRealPower                          float32
	PhaseAApparentPower                      float32
	PhaseBApparentPower                      float32
	PhaseCApparentPower                      float32
	PhaseAReactivePower                      float32
	PhaseBReactivePower                      float32
	PhaseCReactivePower                      float32
	Frequency3                               float32
	PhaseACurrentMagnitude                   float32
	PhaseBCurrentMagnitude                   float32
	PhaseCCurrentMagnitude                   float32
	NeutralCurrentMagnitude                  float32
	PhaseAVoltageMagnitude                   float32
	PhaseBVoltageMagnitude                   float32
	PhaseCVoltageMagnitude                   float32
	PhaseABVoltageMagnitude                  float32
	PhaseBCVoltageMagnitude                  float32
	PhaseCAVoltageMagnitude                  float32
	PhaseACurrentAngle                       float32
	PhaseBCurrentAngle                       float32
	PhaseCCurrentAngle                       float32
	NeutralCurrentAngle                      float32
	PhaseAVoltageAngle                       float32
	PhaseBVoltageAngle                       float32
	PhaseCVoltageAngle                       float32
	PhaseABVoltageAngle                      float32
	PhaseBCVoltageAngle                      float32
	PhaseCAVoltageAngle                      float32
	DisplacementPowerFactor                  float32
	RealPowerDelivered                       float32
	RealPowerReceived                        float32
	ApparentPowerDelivered                   float32
	ApparentPowerReceived                    float32
	ReactivePowerDelivered                   float32
	ReactivePowerReceived                    float32
	NetRealPowerDeliveredReceived            float32
	NetReactivePowerDeliveredReceived        float32
	MVMeterPhaseARMSCurrent                  float32
	MVMeterPhaseBRMSCurrent                  float32
	MVMeterPhaseCRMSCurrent                  float32
	MVMeterNeutralRMSCurrent                 float32
	MVMeterPhaseARMSVoltage                  float32
	MVMeterPhaseBRMSVoltage                  float32
	MVMeterPhaseCRMSVoltage                  float32
	MVMeterPhaseABRMSVoltage                 float32
	MVMeterPhaseBCRMSVoltage                 float32
	MVMeterPhaseCARMSVoltage                 float32
	MVMeterRealPower                         float32
	MVMeterApparentPower                     float32
	MVMeterReactivePower                     float32
	MVMeterPhaseARealPower                   float32
	MVMeterPhaseBRealPower                   float32
	MVMeterPhaseCRealPower                   float32
	MVMeterPhaseAApparentPower               float32
	MVMeterPhaseBApparentPower               float32
	MVMeterPhaseCApparentPower               float32
	MVMeterPhaseAReactivePower               float32
	MVMeterPhaseBReactivePower               float32
	MVMeterPhaseCReactivePower               float32
	MVMeterFrequency                         float32
	MVMeterPhaseACurrentMagnitude            float32
	MVMeterPhaseBCurrentMagnitude            float32
	MVMeterPhaseCCurrentMagnitude            float32
	MVMeterNeutralCurrentMagnitude           float32
	MVMeterPhaseAVoltageMagnitude            float32
	MVMeterPhaseBVoltageMagnitude            float32
	MVMeterPhaseCVoltageMagnitude            float32
	MVMeterPhaseABVoltageMagnitude           float32
	MVMeterPhaseBCVoltageMagnitude           float32
	MVMeterPhaseCAVoltageMagnitude           float32
	MVMeterPhaseACurrentAngle                float32
	MVMeterPhaseBCurrentAngle                float32
	MVMeterPhaseCCurrentAngle                float32
	MVMeterNeutralCurrentAngle               float32
	MVMeterPhaseAVoltageAngle                float32
	MVMeterPhaseBVoltageAngle                float32
	MVMeterPhaseCVoltageAngle                float32
	MVMeterPhaseABVoltageAngle               float32
	MVMeterPhaseBCVoltageAngle               float32
	MVMeterPhaseCAVoltageAngle               float32
	MVMeterDisplacementPowerFactor           float32
	HVMeterRealPowerDelivered                float32
	HVMeterRealPowerReceived                 float32
	HVMeterApparentPowerDelivered            float32
	HVMeterApparentPowerReceived             float32
	HVMeterReactivePowerDelivered            float32
	HVMeterReactivePowerReceived             float32
	HVMeterNetRealPowerDeliveredReceived     float32
	HVMeterNetReactivePowerDeliveredReceived float32
	ATSXfmrPhaseANVoltage                    float32
	ATSXfmrPhaseBNVoltage                    float32
	ATSXfmrPhaseABVoltage                    float32
	ATSGeneratorPhaseANVoltage               float32
	ATSGeneratorPhaseBNVoltage               float32
	ATSGeneratorPhaseABVoltage               float32
	ATSXfmrFrequency                         float32
	ATSGeneratorFrequency                    float32
	NumberOfTransfersForGensetToSupply       float32
	GeneratorFrequency                       float32
	GeneratorVoltageL3L1                     float32
	GeneratorCurrentPhaseL1                  float32
	GeneratorCurrentPhaseL3                  float32
	GeneratorVoltagePhaseL1N                 float32
	GeneratorVoltagePhaseL3N                 float32
	GeneratorFuelLevel                       float32
	GeneratorCoolantTemp                     float32
	GeneratorBatteryVoltage                  float32
	GeneratorEngineSpeed                     float32
	GeneratorEngineRunTime                   float32
	GeneratorStatus                          float32
	GeneratorSuccessStarts                   float32
	GeneratorOilPressure                     float32
}

// DiscreteRegister
// // source annotation to replace is `parquet:"name=@1@@, type=BOOLEAN" json:"@1@@"`
type DiscreteRegister struct {
	H1ClosedStatus                         bool
	H1OpenStatus                           bool
	ClosedStatus89                         bool
	OpenStatus89                           bool
	A86T1A86BF52H1LockoutRelayOperated     bool
	XFMRT1TopOilTempAlarm                  bool
	XFMRT1WindingTempAlarm                 bool
	XFMRT1HighPressureAlarm                bool
	XFMRT1HighOilLevelMainAlarm            bool
	Trip52H1Coil1                          bool
	Trip52M1Coil1                          bool
	RelayFailureAlarmToSEL2240             bool
	TripLORA86T1                           bool
	MechanicalProtectionAlarm              bool
	Differential                           bool
	PhaseTimeOverCurrent                   bool
	GroundTimeOverCurrent                  bool
	VoltsPerHz                             bool
	PhaseUnderVoltage                      bool
	PhaseOverVoltage                       bool
	NeutralOverVoltage                     bool
	TimeSynch                              bool
	FuseFail                               bool
	CircuitBreakerNOContact                bool
	APhaseInvolvedInTheFault               bool
	BPhaseInvolvedInTheFault               bool
	CPhaseInvolvedInTheFault               bool
	BreakerTrip                            bool
	LossOfPotential                        bool
	TimeSourceAccuracy                     bool
	SoftwareAlarms                         bool
	HardwareAlarm                          bool
	BreakerH1Close                         bool
	BreakerH1Open                          bool
	L1CloseStatus89                        bool
	L1OpenStatus89                         bool
	LOROperate                             bool
	XFMRWindingTemperature                 bool
	XFMRLowPressure                        bool
	XFMRLowOilLevel                        bool
	TripBreakerH1                          bool
	TripBreakerM1                          bool
	TripB86T1                              bool
	TripB86BFH1                            bool
	RelayFailAlarm                         bool
	UnderFrequencyProtection               bool
	OverFrequencyProtection                bool
	DifferentialProtection                 bool
	HVWindingTimeOCProtection              bool
	MVWindingTimeOCProtection              bool
	VoltsPerHz2                            bool
	UnderVoltageProtection                 bool
	OverVoltageProtection                  bool
	OverVoltageNeutralProtection           bool
	HVMeterTimeBasedOnIRIGBTimeSource      bool
	HVMeterIRIGBInHighAccuracyMode         bool
	HVMeterSoftwareAlarm                   bool
	HVMeterHardwareAlarm                   bool
	HVMeterLDPFD3LeadLag                   bool
	HVMeterMainBoardOutput3                bool
	MVMeterTimeBasedOnIRIGBTimeSource      bool
	MVMeterIRIGBInHighAccuracyMode         bool
	MVMeterSoftwareAlarm                   bool
	MVMeterHardwareAlarm                   bool
	MVMeterLDPFD3LeadLag                   bool
	MVMeterMainBoardOutput3                bool
	CommonFailAlarm                        bool
	ATSInAutoMode                          bool
	ATSConnectedToTransformer              bool
	A86BF52T1LORTroubleAlarm               bool
	A86T1LORTroubleAlarm                   bool
	B86BF52T1LORTroubleAlarm               bool
	B86T1LORTroubleAlarm                   bool
	A86BF52T1LOROperated                   bool
	A86T1LOROperated                       bool
	B86BF52T1LOROperated                   bool
	B86T1LOROperated                       bool
	GE845Normal                            bool
	SEL487ENormal                          bool
	SEL2730MNormal                         bool
	SEL2401Normal                          bool
	SEL735HVMETERNormal                    bool
	SEL735MVMETERNormal                    bool
	InverterOutputNormal                   bool
	HVACFailAlarm                          bool
	Section1BattChgrAlarm                  bool
	Section3BattChgrAlarm                  bool
	L1Closed89                             bool
	L1Open89                               bool
	L1LossOfACAlarm89                      bool
	L1LossOfDCAlarm89                      bool
	H1Closed52                             bool
	H1Open52                               bool
	H1LowSF6PressAlarm52                   bool
	H1LowSF6PressCutout52                  bool
	XFMRT1TopOilTempAlarm70C               bool
	XFMRT1WindingTempAlarm75C              bool
	XFMRT1LowN2CylinderPressAlarm200PSI    bool
	XFMRT1PressReliefAlarmMainTank         bool
	XFMRT1LossOfVoltageMainSupplyAlarm     bool
	XFMRT1LossOfVoltageFanSupplyAlarm      bool
	XFMRT1LossOfVoltageCoolingControlAlarm bool
	T1Closed89                             bool
	T1Open89                               bool
	T1LossOfACAlarm89                      bool
	T1LossOfDCAlarm89                      bool
}
