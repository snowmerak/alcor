// Automatically generated rust module for 'candidate.proto' file

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
pub struct Request<'a> {
    pub VoterID: Cow<'a, [u8]>,
    pub Signature: Cow<'a, [u8]>,
    pub KeyBytes: Cow<'a, [u8]>,
    pub Name: Cow<'a, str>,
    pub Markdown: Cow<'a, str>,
    pub Hash: Cow<'a, [u8]>,
}

impl<'a> MessageRead<'a> for Request<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.VoterID = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(18) => msg.Signature = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(26) => msg.KeyBytes = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(34) => msg.Name = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(42) => msg.Markdown = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(50) => msg.Hash = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Request<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.VoterID == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.VoterID).len()) }
        + if self.Signature == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Signature).len()) }
        + if self.KeyBytes == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.KeyBytes).len()) }
        + if self.Name == "" { 0 } else { 1 + sizeof_len((&self.Name).len()) }
        + if self.Markdown == "" { 0 } else { 1 + sizeof_len((&self.Markdown).len()) }
        + if self.Hash == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Hash).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.VoterID != Cow::Borrowed(b"") { w.write_with_tag(10, |w| w.write_bytes(&**&self.VoterID))?; }
        if self.Signature != Cow::Borrowed(b"") { w.write_with_tag(18, |w| w.write_bytes(&**&self.Signature))?; }
        if self.KeyBytes != Cow::Borrowed(b"") { w.write_with_tag(26, |w| w.write_bytes(&**&self.KeyBytes))?; }
        if self.Name != "" { w.write_with_tag(34, |w| w.write_string(&**&self.Name))?; }
        if self.Markdown != "" { w.write_with_tag(42, |w| w.write_string(&**&self.Markdown))?; }
        if self.Hash != Cow::Borrowed(b"") { w.write_with_tag(50, |w| w.write_bytes(&**&self.Hash))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct Result_pb<'a> {
    pub OK: bool,
    pub Error: Cow<'a, str>,
}

impl<'a> MessageRead<'a> for Result_pb<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(8) => msg.OK = r.read_bool(bytes)?,
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
        + if self.OK == false { 0 } else { 1 + sizeof_varint(*(&self.OK) as u64) }
        + if self.Error == "" { 0 } else { 1 + sizeof_len((&self.Error).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.OK != false { w.write_with_tag(8, |w| w.write_bool(*&self.OK))?; }
        if self.Error != "" { w.write_with_tag(18, |w| w.write_string(&**&self.Error))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct Candidate<'a> {
    pub Name: Cow<'a, str>,
    pub Markdown: Cow<'a, str>,
    pub Error: Cow<'a, str>,
}

impl<'a> MessageRead<'a> for Candidate<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.Name = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(18) => msg.Markdown = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(26) => msg.Error = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Candidate<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.Name == "" { 0 } else { 1 + sizeof_len((&self.Name).len()) }
        + if self.Markdown == "" { 0 } else { 1 + sizeof_len((&self.Markdown).len()) }
        + if self.Error == "" { 0 } else { 1 + sizeof_len((&self.Error).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.Name != "" { w.write_with_tag(10, |w| w.write_string(&**&self.Name))?; }
        if self.Markdown != "" { w.write_with_tag(18, |w| w.write_string(&**&self.Markdown))?; }
        if self.Error != "" { w.write_with_tag(26, |w| w.write_string(&**&self.Error))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct List<'a> {
    pub Error: Cow<'a, str>,
    pub Names: Vec<Cow<'a, str>>,
}

impl<'a> MessageRead<'a> for List<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.Error = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(18) => msg.Names.push(r.read_string(bytes).map(Cow::Borrowed)?),
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
        + if self.Error == "" { 0 } else { 1 + sizeof_len((&self.Error).len()) }
        + self.Names.iter().map(|s| 1 + sizeof_len((s).len())).sum::<usize>()
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.Error != "" { w.write_with_tag(10, |w| w.write_string(&**&self.Error))?; }
        for s in &self.Names { w.write_with_tag(18, |w| w.write_string(&**s))?; }
        Ok(())
    }
}

