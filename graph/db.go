package graph

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	UserCollection                     *mongo.Collection
	DeviceInstanceCollection           *mongo.Collection
	DeviceModelCollection              *mongo.Collection
	SensorCollection                   *mongo.Collection
	BatchCollection                    *mongo.Collection
	ClientCollection                   *mongo.Collection
	CostCollection                     *mongo.Collection
	SiteCollection                     *mongo.Collection
	GatewayCollection                  *mongo.Collection
	GatewayInstanceCollection          *mongo.Collection
	AlarmConfigurationCollection      *mongo.Collection
	BillMasterCollection               *mongo.Collection
	EmailGatewayCollection             *mongo.Collection
	WorkGroupCollection                *mongo.Collection
	SmartMeterConfigurationCollection *mongo.Collection
)

func ConnectMongo() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connect error:", err)
	}

	// Verify connection
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping error:", err)
	}

	db := client.Database("graphql_demo")
	UserCollection = db.Collection("users")
	DeviceInstanceCollection = db.Collection("device_instance")
	DeviceModelCollection = db.Collection("device_model")
	SensorCollection = db.Collection("sensor")
	BatchCollection = db.Collection("batch")
	ClientCollection = db.Collection("client")
	CostCollection = db.Collection("cost_management_db")
	SiteCollection = db.Collection("industry")
	GatewayCollection = db.Collection("gateway")
	GatewayInstanceCollection = db.Collection("gateway_instance")
	AlarmConfigurationCollection = db.Collection("alarm_configuration")
	BillMasterCollection = db.Collection("bill_master")
	EmailGatewayCollection = db.Collection("email_gateway")
	WorkGroupCollection = db.Collection("work_group")
	SmartMeterConfigurationCollection = db.Collection("smart_meter_configuration")

	SeedData()

	log.Println("MongoDB connected successfully")
}

func SeedData() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Check/Seed User
	count, _ := UserCollection.CountDocuments(ctx, bson.M{})
	if count == 0 {
		UserCollection.InsertOne(ctx, User{
			ID:    "user_1",
			Name:  "John Doe",
			Email: "john.doe@example.com",
		})
		log.Println("Seeded User")
	}

	// 2. Check/Seed Client
	count, _ = ClientCollection.CountDocuments(ctx, bson.M{})
	if count == 0 {
		ClientCollection.InsertOne(ctx, Client{
			ID:                "client_1",
			ClientID:          "c1",
			ClientName:        "Acme Corporation",
			ClientDescription: "Industrial Systems",
			ClientAddress:     "123 Energy Way, Austin, TX",
			ClientLogo:        "https://acme.org/logo.png",
			UserID:            "user_1",
		})
		log.Println("Seeded Client")
	}

	// 3. Check/Seed Site
	count, _ = SiteCollection.CountDocuments(ctx, bson.M{})
	if count == 0 {
		SiteCollection.InsertOne(ctx, Site{
			ID:       "site_1",
			SiteID:   "s1",
			SiteName: "Austin Factory HQ",
			ClientID: "c1",
		})
		log.Println("Seeded Site")
	}

	// 4. Check/Seed Gateway
	count, _ = GatewayCollection.CountDocuments(ctx, bson.M{})
	if count == 0 {
		GatewayCollection.InsertOne(ctx, Gateway{
			ID:          "gateway_1",
			GatewayID:   "g1",
			GatewayName: "HQ Gateway Edge-A",
			SiteID:      "s1",
		})
		log.Println("Seeded Gateway")
	}

	// 5. Check/Seed DeviceInstance
	count, _ = DeviceInstanceCollection.CountDocuments(ctx, bson.M{})
	if count == 0 {
		DeviceInstanceCollection.InsertOne(ctx, DeviceInstance{
			ID:               "device_1",
			DeviceInstanceID: "d1",
			SiteID:           "s1",
			GatewayID:        "g1",
			GeneralInfo: DeviceInstanceInfo{
				DeviceName: "Boiler controller",
			},
		})
		log.Println("Seeded DeviceInstance")
	}

	// 6. Check/Seed Sensor
	count, _ = SensorCollection.CountDocuments(ctx, bson.M{})
	if count == 0 {
		SensorCollection.InsertOne(ctx, Sensor{
			ID:               "sensor_1",
			DeviceInstanceID: "d1",
			SiteID:           "s1",
			GeneralInfo: SensorInfo{
				DeviceModelName: "TempMod-100",
				DeviceName:      "Chamber Temperature Sensor",
				Make:            "Honeywell",
				ModelNumber:     "HW-T100",
			},
		})
		log.Println("Seeded Sensor")
	}
}
