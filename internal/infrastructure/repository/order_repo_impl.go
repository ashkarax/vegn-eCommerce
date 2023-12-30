package repository

import (
	"errors"
	"fmt"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(DB *gorm.DB) interfaceRepository.IOrderRepo {
	return &OrderRepo{DB: DB}
}

func (d *OrderRepo) PlaceNewOrder(orderDetails *requestmodels.OrderDetails) (*string, error) {
	var orderId string

	query := "INSERT INTO orders (user_id, address_id, payment_method, razor_pay_id,order_date,coupon_id) VALUES (?,?,?,?,NOW(),?) RETURNING order_id"
	result := d.DB.Raw(query, orderDetails.UserID, orderDetails.AddressID, orderDetails.PaymentMethod, orderDetails.OrderIdRazorPay, orderDetails.CouponId).Scan(&orderId)
	if result.Error != nil {
		return &orderId, result.Error
	}
	return &orderId, nil
}

func (d *OrderRepo) AddDishesToOrderedItems(orderTableId *string, orderItem *responsemodels.CartItemInfo, OrderDetails *requestmodels.OrderDetails) error {
	fmt.Println("--------------", orderItem.CategoryOfferID)

	query := "INSERT INTO ordered_items (order_id, dish_id, order_quantity, mrp,promotion_discount,dish_price, restaurant_id, order_status, payment_status,category_offer_id,discount_percentage) VALUES (?, ?, ?, ?, ?, ?, ?,?,?,?,?)"

	result := d.DB.Exec(query,
		orderTableId,
		orderItem.DishID,
		orderItem.Quantity,
		orderItem.MRP,
		orderItem.PromotionDiscount,
		orderItem.SalePrice,
		orderItem.RestaurantID,
		OrderDetails.OrderStatus,
		OrderDetails.PaymentStatus,
		orderItem.CategoryOfferID,
		orderItem.DiscountPercentage,
	)
	if result.Error != nil {
		fmt.Println("----------------------", result.Error)
		return result.Error
	}
	return nil

}

func (d *OrderRepo) GetAllOrdersByUser(userId *string) (*[]responsemodels.OrderDetailsResponse, error) {
	var orderSlice []responsemodels.OrderDetailsResponse
	query := "SELECT orders.order_id,ordered_items.dish_id,ordered_items.restaurant_id,ordered_items.ordered_items_id,orders.address_id,orders.payment_method,orders.order_date,ordered_items.order_status,ordered_items.order_quantity,ordered_items.dish_price,restaurants.restaurant_name,dishes.name,dishes.image_url1,addresses.line1,addresses.postal_code,addresses.phone FROM ordered_items JOIN orders ON ordered_items.order_id = orders.order_id JOIN restaurants ON ordered_items.restaurant_id = restaurants.id JOIN dishes ON ordered_items.dish_id = dishes.id JOIN addresses ON orders.address_id = addresses.id WHERE orders.user_id = ?"
	result := d.DB.Raw(query, userId).Scan(&orderSlice)
	if result.RowsAffected == 0 {
		return &orderSlice, errors.New("no past orders by this user")
	}
	if result.Error != nil {
		return &orderSlice, result.Error
	}

	return &orderSlice, nil
}

// func (d *OrderRepo) OrdersForRestaurantById(restaurantId *string) (*[]responsemodels.OrderDetailsResponse, error) {
// 	var orderDetails []responsemodels.OrderDetailsResponse

// 	query := " SELECT orders.order_id,ordered_items.ordered_items_id, ordered_items.dish_id, ordered_items.restaurant_id, orders.address_id,orders.payment_method, orders.order_date, ordered_items.order_status,ordered_items.order_quantity, ordered_items.dish_price,restaurants.restaurant_name, dishes.name, dishes.image_url1,addresses.line1, addresses.street, addresses.city, addresses.postal_code, addresses.phone, addresses.alternate_phone FROM orders JOIN ordered_items ON ordered_items.order_id = orders.order_id JOIN restaurants ON ordered_items.restaurant_id = restaurants.id JOIN dishes ON ordered_items.dish_id = dishes.id JOIN addresses ON orders.address_id = addresses.id WHERE ordered_items.restaurant_id = ? AND ordered_items.order_status='processing';"
// 	result := d.DB.Raw(query, restaurantId).Scan(&orderDetails)
// 	if result.RowsAffected == 0 {
// 		return &orderDetails, errors.New("no orders for your restaurant")
// 	}
// 	if result.Error != nil {
// 		return &orderDetails, result.Error
// 	}

// 	return &orderDetails, nil

// }
func (d *OrderRepo) OrdersForRestaurantById(restaurantId *string) (*[]responsemodels.OrderDataFetcherX, error) {
	var orderDetails []responsemodels.OrderDataFetcherX

	query := " SELECT orders.order_id,orders.payment_method,orders.order_date,	addresses.line1,addresses.street,addresses.city,addresses.postal_code,users.f_name,users.l_name,ordered_items.ordered_items_id,ordered_items.payment_status,ordered_items.dish_id,ordered_items.order_quantity,ordered_items.dish_price,ordered_items.order_status,dishes.name,dishes.image_url1,dishes.portion_size  FROM orders INNER JOIN ordered_items ON orders.order_id = ordered_items.order_id INNER JOIN dishes ON ordered_items.dish_id = dishes.id INNER JOIN	addresses ON orders.address_id = addresses.id INNER JOIN users ON orders.user_id = users.id  WHERE	ordered_items.restaurant_id = ? ORDER BY orders.order_date DESC; "
	result := d.DB.Raw(query, restaurantId).Scan(&orderDetails)
	if result.RowsAffected == 0 {
		return &orderDetails, errors.New("no orders for your restaurant")
	}
	if result.Error != nil {
		return &orderDetails, result.Error
	}

	return &orderDetails, nil

}

func (d *OrderRepo) GetOrderDetailsByOrderId(userId *string, orderId *string) (*[]responsemodels.RazorpayResponse, error) {
	var orderDetails []responsemodels.RazorpayResponse

	query := "SELECT orders.razor_pay_id,ordered_items.order_quantity,ordered_items.dish_price,users.f_name,users.l_name,users.email,users.phone	FROM orders	INNER JOIN ordered_items ON orders.order_id = ordered_items.order_id	INNER JOIN users ON orders.user_id = users.id WHERE users.id=? AND orders.order_id=?"
	result := d.DB.Raw(query, userId, orderId).Scan(&orderDetails)
	if result.RowsAffected == 0 {
		return &orderDetails, errors.New("can't find order,input a valid order-id")
	}
	if result.Error != nil {
		return &orderDetails, result.Error
	}

	return &orderDetails, nil
}

func (d *OrderRepo) UpdateStatusToSuccess(userId *string, orderId *string, razorPayId *string) (*[]responsemodels.OrderDetailsResponse, error) {
	var orderDetails []responsemodels.OrderDetailsResponse

	query := "UPDATE ordered_items SET payment_status='success', order_status='processing' FROM orders WHERE ordered_items.order_id=orders.order_id AND orders.user_id=? AND orders.order_id=? AND orders.razor_pay_id=? RETURNING*"
	result := d.DB.Raw(query, userId, orderId, razorPayId).Scan(&orderDetails)
	if result.Error != nil {
		return nil, errors.New("face some issue while update order status and payment status on verify online payment success")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("rows affected is zero,some error happened")
	}
	return &orderDetails, nil

}

func (d *OrderRepo) ReturnOrderStats(orderData *requestmodels.CanOrRetReq) (*responsemodels.CanOrRetResp, error) {
	var orderDataResp responsemodels.CanOrRetResp

	query := "SELECT orders.payment_method,ordered_items.payment_status,ordered_items.order_status,ordered_items.order_quantity,ordered_items.dish_price,ordered_items.dish_id,ordered_items.restaurant_id,ordered_items.deliver_date FROM orders JOIN ordered_items  ON orders.order_id = ordered_items.order_id WHERE orders.user_id = ?    AND ordered_items.ordered_items_id = ?"
	result := d.DB.Raw(query, orderData.UserID, orderData.OrderedItemsID).Scan(&orderDataResp)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("rows affected is zero,enter valid orderid/dishid")
	}
	return &orderDataResp, nil
}

func (d *OrderRepo) ChangeOrderStatus(orderedItemsID *string, newStatus *string) error {
	query := "UPDATE ordered_items SET order_status=? WHERE ordered_items_id=?"
	err := d.DB.Exec(query, newStatus, orderedItemsID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *OrderRepo) UpdatePaymentStatus(orderedItemsID *string, newStatus *string) error {
	query := "UPDATE ordered_items SET payment_status=? WHERE ordered_items_id=?"
	err := d.DB.Exec(query, newStatus, orderedItemsID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *OrderRepo) RestReturnOrderStatus(restaurantId *string, ordereditemsid *string) (*responsemodels.CanOrRetResp, error) {
	var orderDataResp responsemodels.CanOrRetResp

	query := "SELECT orders.payment_method,ordered_items.payment_status,ordered_items.order_status,ordered_items.order_quantity,ordered_items.dish_price,ordered_items.dish_id,ordered_items.restaurant_id,ordered_items.deliver_date FROM orders JOIN ordered_items  ON orders.order_id = ordered_items.order_id WHERE ordered_items.restaurant_id = ?    AND ordered_items.ordered_items_id = ?"
	result := d.DB.Raw(query, restaurantId, ordereditemsid).Scan(&orderDataResp)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("rows affected is zero,enter valid restaurantid/dishid")
	}
	return &orderDataResp, nil
}

func (d *OrderRepo) UpdateDeliveryDate(ordereditemsid *string) {
	query := "UPDATE ordered_items SET deliver_date= NOW() WHERE ordered_items_id = ?"
	r := d.DB.Exec(query, ordereditemsid).Error

	if r != nil {
		fmt.Println("---------------", r)
	}

}

func (d *OrderRepo) GetOrderDataForPDFByIds(userId *string, orderId *string) (*responsemodels.OrderDetailsPDF, error) {
	var orderDetails responsemodels.OrderDetailsPDF

	query := "SELECT orders.*, users.*, coupons.*,addresses.* FROM orders INNER JOIN users ON orders.user_id = users.id INNER JOIN addresses ON orders.address_id = addresses.id LEFT JOIN coupons ON orders.coupon_id = coupons.coupon_id WHERE users.id = ? AND orders.order_id = ?;"
	result := d.DB.Raw(query, userId, orderId).Scan(&orderDetails)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {

		return nil, errors.New("can't find your order,orderid not matching by the user")
	}
	return &orderDetails, nil

}

func (d *OrderRepo) GetOrderDetailsForPDFById(orderId *string) (*[]responsemodels.OrderedItemsDataPDF, error) {
	var orderItemsDetails []responsemodels.OrderedItemsDataPDF

	query := "SELECT ordered_items.order_quantity,restaurants.restaurant_name,dishes.name,dishes.mrp,dishes.promotion_discount,category_offers.discount_percentage 	FROM ordered_items INNER JOIN restaurants ON ordered_items.restaurant_id=restaurants.id INNER JOIN dishes ON ordered_items.dish_id=dishes.id LEFT JOIN category_offers ON ordered_items.category_offer_id=category_offers.category_offer_id WHERE ordered_items.order_id=?;"
	result := d.DB.Raw(query, orderId).Scan(&orderItemsDetails).Error
	if result != nil {
		return &orderItemsDetails, result
	}

	return &orderItemsDetails, nil

}

func (d *OrderRepo) OrdersForSalesReportByRestId(restId *string) (*[]responsemodels.OrderDetaisForSalesReport, error) {
	var orderDetails []responsemodels.OrderDetaisForSalesReport

	query := "SELECT orders.order_id,orders.payment_method,orders.order_date,ordered_items.dish_id,ordered_items.order_quantity,ordered_items.mrp,ordered_items.promotion_discount,ordered_items.discount_percentage,ordered_items.dish_price,ordered_items.deliver_date, dishes.name FROM orders INNER JOIN ordered_items ON orders.order_id = ordered_items.order_id INNER JOIN dishes ON ordered_items.dish_id = dishes.id WHERE ordered_items.order_status = 'delivered' AND ordered_items.restaurant_id = ?"
	result := d.DB.Raw(query, restId).Scan(&orderDetails)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {

		return nil, errors.New("no order records found for your restaurant")
	}

	return &orderDetails, nil

}

func (d *OrderRepo) GetDataSalesReportForCustomDays(restaurantId *string, customDays *string) (*responsemodels.SalesReportData, error) {
	var returnData responsemodels.SalesReportData

	remainingQuery := "(now() - interval '" + *customDays + " day')"
	query := "SELECT COUNT(*) AS Orders, SUM(ordered_items.order_quantity) AS Quantity, SUM(ordered_items.dish_price) AS Price FROM orders INNER JOIN ordered_items ON orders.order_id=ordered_items.order_id  WHERE ordered_items.restaurant_id = ? AND ordered_items.order_status='delivered' AND orders.order_date >= " + remainingQuery
	result := d.DB.Raw(query, restaurantId).Scan(&returnData)

	if result.Error != nil {
		return nil, errors.New("face some issue while get report by days")
	}

	return &returnData, nil

}

func (d *OrderRepo) GetDataSalesReportYYMMDD(restId *string, yymmdd *requestmodels.SalesReportYYMMDD) (*responsemodels.SalesReportData, error) {
	var returnData responsemodels.SalesReportData

	var remainingQuery string

	if yymmdd.Year != "" {
		remainingQuery = " EXTRACT(YEAR FROM order_date)=" + yymmdd.Year
	}
	if yymmdd.Year != "" && yymmdd.Month != "" {
		remainingQuery = " EXTRACT(YEAR FROM order_date)=" + yymmdd.Year + " AND EXTRACT(Month FROM order_date)=" + yymmdd.Month
	}
	if yymmdd.Year != "" && yymmdd.Month != "" && yymmdd.Day != "" {
		remainingQuery = " EXTRACT(YEAR FROM order_date)=" + yymmdd.Year + " AND EXTRACT(Month FROM order_date)=" + yymmdd.Month + " AND EXTRACT(Day FROM order_date)=" + yymmdd.Day
	}

	query := "SELECT COUNT(*) AS Orders, SUM(ordered_items.order_quantity) AS Quantity, SUM(ordered_items.dish_price) AS Price FROM orders INNER JOIN ordered_items ON orders.order_id=ordered_items.order_id  WHERE ordered_items.restaurant_id = ? AND ordered_items.order_status='delivered'  AND" + remainingQuery
	result := d.DB.Raw(query, restId).Scan(&returnData)
	if result.Error != nil {
		return nil, result.Error
	}

	return &returnData, nil

}
