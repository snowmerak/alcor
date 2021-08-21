// Automatically generated rust module for 'stats.proto' file

#![allow(non_snake_case)]
#![allow(non_upper_case_globals)]
#![allow(non_camel_case_types)]
#![allow(unused_imports)]
#![allow(unknown_lints)]
#![allow(clippy::all)]
#![cfg_attr(rustfmt, rustfmt_skip)]


use std::borrow::Cow;
use quick_protobuf::{MessageRead, MessageWrite, BytesReader, Writer, WriterBackend, Result};
use quick_protobuf::sizeofs::*;
use super::*;

#[derive(Debug, Default, PartialEq, Clone)]
pub struct Data<'a> {
    pub Candidate: Cow<'a, str>,
    pub Gender: bool,
    pub Age: i32,
    pub Region: Cow<'a, str>,
    pub Job: Cow<'a, str>,
    pub Education: Cow<'a, str>,
    pub Married: bool,
    pub Divorced: bool,
    pub HasCar: bool,
    pub HouseType: Cow<'a, str>,
    pub Salary: i32,
    pub HasDebt: bool,
    pub Ideology: Cow<'a, str>,
    pub VoterID: Cow<'a, [u8]>,
}

impl<'a> MessageRead<'a> for Data<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.Candidate = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(16) => msg.Gender = r.read_bool(bytes)?,
                Ok(24) => msg.Age = r.read_int32(bytes)?,
                Ok(34) => msg.Region = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(42) => msg.Job = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(50) => msg.Education = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(56) => msg.Married = r.read_bool(bytes)?,
                Ok(64) => msg.Divorced = r.read_bool(bytes)?,
                Ok(72) => msg.HasCar = r.read_bool(bytes)?,
                Ok(82) => msg.HouseType = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(88) => msg.Salary = r.read_int32(bytes)?,
                Ok(96) => msg.HasDebt = r.read_bool(bytes)?,
                Ok(106) => msg.Ideology = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(114) => msg.VoterID = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Data<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.Candidate == "" { 0 } else { 1 + sizeof_len((&self.Candidate).len()) }
        + if self.Gender == false { 0 } else { 1 + sizeof_varint(*(&self.Gender) as u64) }
        + if self.Age == 0i32 { 0 } else { 1 + sizeof_varint(*(&self.Age) as u64) }
        + if self.Region == "" { 0 } else { 1 + sizeof_len((&self.Region).len()) }
        + if self.Job == "" { 0 } else { 1 + sizeof_len((&self.Job).len()) }
        + if self.Education == "" { 0 } else { 1 + sizeof_len((&self.Education).len()) }
        + if self.Married == false { 0 } else { 1 + sizeof_varint(*(&self.Married) as u64) }
        + if self.Divorced == false { 0 } else { 1 + sizeof_varint(*(&self.Divorced) as u64) }
        + if self.HasCar == false { 0 } else { 1 + sizeof_varint(*(&self.HasCar) as u64) }
        + if self.HouseType == "" { 0 } else { 1 + sizeof_len((&self.HouseType).len()) }
        + if self.Salary == 0i32 { 0 } else { 1 + sizeof_varint(*(&self.Salary) as u64) }
        + if self.HasDebt == false { 0 } else { 1 + sizeof_varint(*(&self.HasDebt) as u64) }
        + if self.Ideology == "" { 0 } else { 1 + sizeof_len((&self.Ideology).len()) }
        + if self.VoterID == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.VoterID).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.Candidate != "" { w.write_with_tag(10, |w| w.write_string(&**&self.Candidate))?; }
        if self.Gender != false { w.write_with_tag(16, |w| w.write_bool(*&self.Gender))?; }
        if self.Age != 0i32 { w.write_with_tag(24, |w| w.write_int32(*&self.Age))?; }
        if self.Region != "" { w.write_with_tag(34, |w| w.write_string(&**&self.Region))?; }
        if self.Job != "" { w.write_with_tag(42, |w| w.write_string(&**&self.Job))?; }
        if self.Education != "" { w.write_with_tag(50, |w| w.write_string(&**&self.Education))?; }
        if self.Married != false { w.write_with_tag(56, |w| w.write_bool(*&self.Married))?; }
        if self.Divorced != false { w.write_with_tag(64, |w| w.write_bool(*&self.Divorced))?; }
        if self.HasCar != false { w.write_with_tag(72, |w| w.write_bool(*&self.HasCar))?; }
        if self.HouseType != "" { w.write_with_tag(82, |w| w.write_string(&**&self.HouseType))?; }
        if self.Salary != 0i32 { w.write_with_tag(88, |w| w.write_int32(*&self.Salary))?; }
        if self.HasDebt != false { w.write_with_tag(96, |w| w.write_bool(*&self.HasDebt))?; }
        if self.Ideology != "" { w.write_with_tag(106, |w| w.write_string(&**&self.Ideology))?; }
        if self.VoterID != Cow::Borrowed(b"") { w.write_with_tag(114, |w| w.write_bytes(&**&self.VoterID))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct List<'a> {
    pub Data: Vec<Cow<'a, str>>,
}

impl<'a> MessageRead<'a> for List<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.Data.push(r.read_string(bytes).map(Cow::Borrowed)?),
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for List<'a> {
    fn get_size(&self) -> usize {
        0
        + self.Data.iter().map(|s| 1 + sizeof_len((s).len())).sum::<usize>()
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        for s in &self.Data { w.write_with_tag(10, |w| w.write_string(&**s))?; }
        Ok(())
    }
}

