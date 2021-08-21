// Automatically generated rust module for 'info.proto' file

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
pub struct Voter<'a> {
    pub HashID: Cow<'a, [u8]>,
    pub PublicKey: Cow<'a, [u8]>,
    pub RandomBytes: Cow<'a, [u8]>,
    pub Timestamp: Cow<'a, [u8]>,
    pub Voted: bool,
    pub IsCandidate: bool,
    pub RepliedSurvey: bool,
}

impl<'a> MessageRead<'a> for Voter<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.HashID = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(18) => msg.PublicKey = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(26) => msg.RandomBytes = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(34) => msg.Timestamp = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(40) => msg.Voted = r.read_bool(bytes)?,
                Ok(48) => msg.IsCandidate = r.read_bool(bytes)?,
                Ok(56) => msg.RepliedSurvey = r.read_bool(bytes)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Voter<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.HashID == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.HashID).len()) }
        + if self.PublicKey == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.PublicKey).len()) }
        + if self.RandomBytes == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.RandomBytes).len()) }
        + if self.Timestamp == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Timestamp).len()) }
        + if self.Voted == false { 0 } else { 1 + sizeof_varint(*(&self.Voted) as u64) }
        + if self.IsCandidate == false { 0 } else { 1 + sizeof_varint(*(&self.IsCandidate) as u64) }
        + if self.RepliedSurvey == false { 0 } else { 1 + sizeof_varint(*(&self.RepliedSurvey) as u64) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.HashID != Cow::Borrowed(b"") { w.write_with_tag(10, |w| w.write_bytes(&**&self.HashID))?; }
        if self.PublicKey != Cow::Borrowed(b"") { w.write_with_tag(18, |w| w.write_bytes(&**&self.PublicKey))?; }
        if self.RandomBytes != Cow::Borrowed(b"") { w.write_with_tag(26, |w| w.write_bytes(&**&self.RandomBytes))?; }
        if self.Timestamp != Cow::Borrowed(b"") { w.write_with_tag(34, |w| w.write_bytes(&**&self.Timestamp))?; }
        if self.Voted != false { w.write_with_tag(40, |w| w.write_bool(*&self.Voted))?; }
        if self.IsCandidate != false { w.write_with_tag(48, |w| w.write_bool(*&self.IsCandidate))?; }
        if self.RepliedSurvey != false { w.write_with_tag(56, |w| w.write_bool(*&self.RepliedSurvey))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct Paper<'a> {
    pub VoterID: Cow<'a, [u8]>,
    pub Timestamp: Cow<'a, [u8]>,
    pub Message: Cow<'a, [u8]>,
    pub RandomBytes: Cow<'a, [u8]>,
    pub Hash: Cow<'a, [u8]>,
    pub Signature: Cow<'a, [u8]>,
}

impl<'a> MessageRead<'a> for Paper<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.VoterID = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(18) => msg.Timestamp = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(26) => msg.Message = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(34) => msg.RandomBytes = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(42) => msg.Hash = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(50) => msg.Signature = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Paper<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.VoterID == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.VoterID).len()) }
        + if self.Timestamp == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Timestamp).len()) }
        + if self.Message == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Message).len()) }
        + if self.RandomBytes == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.RandomBytes).len()) }
        + if self.Hash == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Hash).len()) }
        + if self.Signature == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Signature).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.VoterID != Cow::Borrowed(b"") { w.write_with_tag(10, |w| w.write_bytes(&**&self.VoterID))?; }
        if self.Timestamp != Cow::Borrowed(b"") { w.write_with_tag(18, |w| w.write_bytes(&**&self.Timestamp))?; }
        if self.Message != Cow::Borrowed(b"") { w.write_with_tag(26, |w| w.write_bytes(&**&self.Message))?; }
        if self.RandomBytes != Cow::Borrowed(b"") { w.write_with_tag(34, |w| w.write_bytes(&**&self.RandomBytes))?; }
        if self.Hash != Cow::Borrowed(b"") { w.write_with_tag(42, |w| w.write_bytes(&**&self.Hash))?; }
        if self.Signature != Cow::Borrowed(b"") { w.write_with_tag(50, |w| w.write_bytes(&**&self.Signature))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct Bundle<'a> {
    pub Hash: Cow<'a, [u8]>,
    pub Prev: Cow<'a, [u8]>,
    pub SubHashes: Vec<Cow<'a, [u8]>>,
}

impl<'a> MessageRead<'a> for Bundle<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.Hash = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(18) => msg.Prev = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(26) => msg.SubHashes.push(r.read_bytes(bytes).map(Cow::Borrowed)?),
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Bundle<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.Hash == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Hash).len()) }
        + if self.Prev == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Prev).len()) }
        + self.SubHashes.iter().map(|s| 1 + sizeof_len((s).len())).sum::<usize>()
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.Hash != Cow::Borrowed(b"") { w.write_with_tag(10, |w| w.write_bytes(&**&self.Hash))?; }
        if self.Prev != Cow::Borrowed(b"") { w.write_with_tag(18, |w| w.write_bytes(&**&self.Prev))?; }
        for s in &self.SubHashes { w.write_with_tag(26, |w| w.write_bytes(&**s))?; }
        Ok(())
    }
}

