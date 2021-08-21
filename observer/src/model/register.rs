// Automatically generated rust module for 'register.proto' file

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
pub struct Phone<'a> {
    pub Number: Cow<'a, str>,
    pub Name: Cow<'a, str>,
    pub RRN: Cow<'a, str>,
    pub Gender: bool,
}

impl<'a> MessageRead<'a> for Phone<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(18) => msg.Number = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(26) => msg.Name = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(34) => msg.RRN = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(8) => msg.Gender = r.read_bool(bytes)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Phone<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.Number == "" { 0 } else { 1 + sizeof_len((&self.Number).len()) }
        + if self.Name == "" { 0 } else { 1 + sizeof_len((&self.Name).len()) }
        + if self.RRN == "" { 0 } else { 1 + sizeof_len((&self.RRN).len()) }
        + if self.Gender == false { 0 } else { 1 + sizeof_varint(*(&self.Gender) as u64) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.Number != "" { w.write_with_tag(18, |w| w.write_string(&**&self.Number))?; }
        if self.Name != "" { w.write_with_tag(26, |w| w.write_string(&**&self.Name))?; }
        if self.RRN != "" { w.write_with_tag(34, |w| w.write_string(&**&self.RRN))?; }
        if self.Gender != false { w.write_with_tag(8, |w| w.write_bool(*&self.Gender))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct Cache<'a> {
    pub Number: Cow<'a, str>,
    pub Name: Cow<'a, str>,
    pub RRN: Cow<'a, str>,
    pub Gender: bool,
    pub Code: Cow<'a, [u8]>,
}

impl<'a> MessageRead<'a> for Cache<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(18) => msg.Number = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(26) => msg.Name = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(34) => msg.RRN = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(8) => msg.Gender = r.read_bool(bytes)?,
                Ok(42) => msg.Code = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Cache<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.Number == "" { 0 } else { 1 + sizeof_len((&self.Number).len()) }
        + if self.Name == "" { 0 } else { 1 + sizeof_len((&self.Name).len()) }
        + if self.RRN == "" { 0 } else { 1 + sizeof_len((&self.RRN).len()) }
        + if self.Gender == false { 0 } else { 1 + sizeof_varint(*(&self.Gender) as u64) }
        + if self.Code == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Code).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.Number != "" { w.write_with_tag(18, |w| w.write_string(&**&self.Number))?; }
        if self.Name != "" { w.write_with_tag(26, |w| w.write_string(&**&self.Name))?; }
        if self.RRN != "" { w.write_with_tag(34, |w| w.write_string(&**&self.RRN))?; }
        if self.Gender != false { w.write_with_tag(8, |w| w.write_bool(*&self.Gender))?; }
        if self.Code != Cow::Borrowed(b"") { w.write_with_tag(42, |w| w.write_bytes(&**&self.Code))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct CertCode<'a> {
    pub Code: Cow<'a, [u8]>,
    pub PublicKey: Cow<'a, [u8]>,
}

impl<'a> MessageRead<'a> for CertCode<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.Code = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(18) => msg.PublicKey = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for CertCode<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.Code == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Code).len()) }
        + if self.PublicKey == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.PublicKey).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.Code != Cow::Borrowed(b"") { w.write_with_tag(10, |w| w.write_bytes(&**&self.Code))?; }
        if self.PublicKey != Cow::Borrowed(b"") { w.write_with_tag(18, |w| w.write_bytes(&**&self.PublicKey))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct Result_pb<'a> {
    pub ID: Cow<'a, [u8]>,
    pub Error: Cow<'a, str>,
}

impl<'a> MessageRead<'a> for Result_pb<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.ID = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(18) => msg.Error = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Result_pb<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.ID == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.ID).len()) }
        + if self.Error == "" { 0 } else { 1 + sizeof_len((&self.Error).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.ID != Cow::Borrowed(b"") { w.write_with_tag(10, |w| w.write_bytes(&**&self.ID))?; }
        if self.Error != "" { w.write_with_tag(18, |w| w.write_string(&**&self.Error))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct Account<'a> {
    pub HashID: Cow<'a, [u8]>,
    pub PrivateKey: Cow<'a, [u8]>,
}

impl<'a> MessageRead<'a> for Account<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.HashID = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(18) => msg.PrivateKey = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Account<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.HashID == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.HashID).len()) }
        + if self.PrivateKey == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.PrivateKey).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.HashID != Cow::Borrowed(b"") { w.write_with_tag(10, |w| w.write_bytes(&**&self.HashID))?; }
        if self.PrivateKey != Cow::Borrowed(b"") { w.write_with_tag(18, |w| w.write_bytes(&**&self.PrivateKey))?; }
        Ok(())
    }
}

