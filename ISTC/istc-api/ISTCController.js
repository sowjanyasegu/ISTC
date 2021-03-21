var express = require("express");
var router = express.Router();
var bodyParser = require("body-parser");

router.use(bodyParser.urlencoded({ extended: true }));
router.use(bodyParser.json());

const txSubmit = require("./invoke");
const txFetch = require("./query");

//var TFBC = require("./FabricHelper");

// CreateTea
router.post("/createTeaLot", async function(req, res) {
  try {
    let result = await txSubmit.invoke("createTeaLot", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

// UpdateTea
router.post("/updateTeaLot", async function(req, res) {
  try {
    let result = await txSubmit.invoke("updateTeaLot", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

// createTeaPacket
router.post("/createTeaPacket", async function(req, res) {
  try {
    let result = await txSubmit.invoke("createTeaPacket", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

// createTeaPacket
router.post("/updateTeaPacket", async function(req, res) {
  try {
    let result = await txSubmit.invoke("updateTeaPacket", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

//initiateTeaTaste
router.post("/initiateTeaTaste", async function(req, res) {
  try {
    let result = await txSubmit.invoke("initiateTeaTaste",[req.body.teaLotID,req.body.teaTaster]);
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

//UpdateTeaTaste
router.post("/updateTeaTaste", async function(req, res) {
  try {
    let result = await txSubmit.invoke("updateTeaTaste", [req.body.teaLotID,req.body.teaTasteNotes]);
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

//initiateLabTest
router.post("/initiateLabTest", async function(req, res) {
  try {
    console.log("body : ",req.body);
    let result = await txSubmit.invoke("initiateLabTest", [req.body.teaLotID,req.body.labName]);
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

//updateLabResult
router.post("/updateLabResult", async function(req, res) {
  try {
    let result = await txSubmit.invoke("updateLabResult",[req.body.teaLotID,req.body.certified]);
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

// getTea
router.post("/getTeaLot", async function(req, res) {
  try {
    let result = await txFetch.query("getTeaLot", req.body.teaLotID);
    res.send(JSON.parse(result));
  } catch (err) {
    res.status(500).send(err);
  }
});

// getTeaHistory
router.post("/getTeaLotHistory", async function(req, res) {
  try {
    let result = await txFetch.query("getTeaLotHistory", req.body.teaLotID);
    res.send(JSON.parse(result));
  } catch (err) {
    res.status(500).send(err);
  }
});

// getTea
router.post("/getTeaPacket", async function(req, res) {
  try {
    let result = await txFetch.query("getTeaPacket", req.body.teaPacketID);
    res.send(JSON.parse(result));
  } catch (err) {
    res.status(500).send(err);
  }
});

// getTeaPacketHistory
router.post("/getTeaPacketHistory", async function(req, res) {
  try {
    let result = await txFetch.query("getTeaPacketHistory", req.body.teaPacketID);
    res.send(JSON.parse(result));
  } catch (err) {
    res.status(500).send(err);
  }
});

// queryAllTeaLots
router.get("/queryAllTeaLots", async function(req, res) {
  try {
    let result = await txFetch.query("queryAllTeaLots");
    res.send(JSON.parse(result));
  } catch (err) {
    res.status(500).send(err);
  }
});

// queryAllTeaPackets
router.get("/queryAllTeaPackets", async function(req, res) {
  try {
    let result = await txFetch.query("queryAllTeaPackets");
    res.send(JSON.parse(result));
  } catch (err) {
    res.status(500).send(err);
  }
});


module.exports = router;
