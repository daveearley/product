// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("AccountUsers", testAccountUsers)
	t.Run("Accounts", testAccounts)
	t.Run("Attendees", testAttendees)
	t.Run("Attributes", testAttributes)
	t.Run("Customers", testCustomers)
	t.Run("DiscountCodes", testDiscountCodes)
	t.Run("Events", testEvents)
	t.Run("QuestionAnswers", testQuestionAnswers)
	t.Run("QuestionOptions", testQuestionOptions)
	t.Run("Questions", testQuestions)
	t.Run("TicketReservations", testTicketReservations)
	t.Run("Tickets", testTickets)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodes)
	t.Run("TransactionItems", testTransactionItems)
	t.Run("Transactions", testTransactions)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersDelete)
	t.Run("Accounts", testAccountsDelete)
	t.Run("Attendees", testAttendeesDelete)
	t.Run("Attributes", testAttributesDelete)
	t.Run("Customers", testCustomersDelete)
	t.Run("DiscountCodes", testDiscountCodesDelete)
	t.Run("Events", testEventsDelete)
	t.Run("QuestionAnswers", testQuestionAnswersDelete)
	t.Run("QuestionOptions", testQuestionOptionsDelete)
	t.Run("Questions", testQuestionsDelete)
	t.Run("TicketReservations", testTicketReservationsDelete)
	t.Run("Tickets", testTicketsDelete)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesDelete)
	t.Run("TransactionItems", testTransactionItemsDelete)
	t.Run("Transactions", testTransactionsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersQueryDeleteAll)
	t.Run("Accounts", testAccountsQueryDeleteAll)
	t.Run("Attendees", testAttendeesQueryDeleteAll)
	t.Run("Attributes", testAttributesQueryDeleteAll)
	t.Run("Customers", testCustomersQueryDeleteAll)
	t.Run("DiscountCodes", testDiscountCodesQueryDeleteAll)
	t.Run("Events", testEventsQueryDeleteAll)
	t.Run("QuestionAnswers", testQuestionAnswersQueryDeleteAll)
	t.Run("QuestionOptions", testQuestionOptionsQueryDeleteAll)
	t.Run("Questions", testQuestionsQueryDeleteAll)
	t.Run("TicketReservations", testTicketReservationsQueryDeleteAll)
	t.Run("Tickets", testTicketsQueryDeleteAll)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesQueryDeleteAll)
	t.Run("TransactionItems", testTransactionItemsQueryDeleteAll)
	t.Run("Transactions", testTransactionsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersSliceDeleteAll)
	t.Run("Accounts", testAccountsSliceDeleteAll)
	t.Run("Attendees", testAttendeesSliceDeleteAll)
	t.Run("Attributes", testAttributesSliceDeleteAll)
	t.Run("Customers", testCustomersSliceDeleteAll)
	t.Run("DiscountCodes", testDiscountCodesSliceDeleteAll)
	t.Run("Events", testEventsSliceDeleteAll)
	t.Run("QuestionAnswers", testQuestionAnswersSliceDeleteAll)
	t.Run("QuestionOptions", testQuestionOptionsSliceDeleteAll)
	t.Run("Questions", testQuestionsSliceDeleteAll)
	t.Run("TicketReservations", testTicketReservationsSliceDeleteAll)
	t.Run("Tickets", testTicketsSliceDeleteAll)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesSliceDeleteAll)
	t.Run("TransactionItems", testTransactionItemsSliceDeleteAll)
	t.Run("Transactions", testTransactionsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersExists)
	t.Run("Accounts", testAccountsExists)
	t.Run("Attendees", testAttendeesExists)
	t.Run("Attributes", testAttributesExists)
	t.Run("Customers", testCustomersExists)
	t.Run("DiscountCodes", testDiscountCodesExists)
	t.Run("Events", testEventsExists)
	t.Run("QuestionAnswers", testQuestionAnswersExists)
	t.Run("QuestionOptions", testQuestionOptionsExists)
	t.Run("Questions", testQuestionsExists)
	t.Run("TicketReservations", testTicketReservationsExists)
	t.Run("Tickets", testTicketsExists)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesExists)
	t.Run("TransactionItems", testTransactionItemsExists)
	t.Run("Transactions", testTransactionsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersFind)
	t.Run("Accounts", testAccountsFind)
	t.Run("Attendees", testAttendeesFind)
	t.Run("Attributes", testAttributesFind)
	t.Run("Customers", testCustomersFind)
	t.Run("DiscountCodes", testDiscountCodesFind)
	t.Run("Events", testEventsFind)
	t.Run("QuestionAnswers", testQuestionAnswersFind)
	t.Run("QuestionOptions", testQuestionOptionsFind)
	t.Run("Questions", testQuestionsFind)
	t.Run("TicketReservations", testTicketReservationsFind)
	t.Run("Tickets", testTicketsFind)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesFind)
	t.Run("TransactionItems", testTransactionItemsFind)
	t.Run("Transactions", testTransactionsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersBind)
	t.Run("Accounts", testAccountsBind)
	t.Run("Attendees", testAttendeesBind)
	t.Run("Attributes", testAttributesBind)
	t.Run("Customers", testCustomersBind)
	t.Run("DiscountCodes", testDiscountCodesBind)
	t.Run("Events", testEventsBind)
	t.Run("QuestionAnswers", testQuestionAnswersBind)
	t.Run("QuestionOptions", testQuestionOptionsBind)
	t.Run("Questions", testQuestionsBind)
	t.Run("TicketReservations", testTicketReservationsBind)
	t.Run("Tickets", testTicketsBind)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesBind)
	t.Run("TransactionItems", testTransactionItemsBind)
	t.Run("Transactions", testTransactionsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersOne)
	t.Run("Accounts", testAccountsOne)
	t.Run("Attendees", testAttendeesOne)
	t.Run("Attributes", testAttributesOne)
	t.Run("Customers", testCustomersOne)
	t.Run("DiscountCodes", testDiscountCodesOne)
	t.Run("Events", testEventsOne)
	t.Run("QuestionAnswers", testQuestionAnswersOne)
	t.Run("QuestionOptions", testQuestionOptionsOne)
	t.Run("Questions", testQuestionsOne)
	t.Run("TicketReservations", testTicketReservationsOne)
	t.Run("Tickets", testTicketsOne)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesOne)
	t.Run("TransactionItems", testTransactionItemsOne)
	t.Run("Transactions", testTransactionsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersAll)
	t.Run("Accounts", testAccountsAll)
	t.Run("Attendees", testAttendeesAll)
	t.Run("Attributes", testAttributesAll)
	t.Run("Customers", testCustomersAll)
	t.Run("DiscountCodes", testDiscountCodesAll)
	t.Run("Events", testEventsAll)
	t.Run("QuestionAnswers", testQuestionAnswersAll)
	t.Run("QuestionOptions", testQuestionOptionsAll)
	t.Run("Questions", testQuestionsAll)
	t.Run("TicketReservations", testTicketReservationsAll)
	t.Run("Tickets", testTicketsAll)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesAll)
	t.Run("TransactionItems", testTransactionItemsAll)
	t.Run("Transactions", testTransactionsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersCount)
	t.Run("Accounts", testAccountsCount)
	t.Run("Attendees", testAttendeesCount)
	t.Run("Attributes", testAttributesCount)
	t.Run("Customers", testCustomersCount)
	t.Run("DiscountCodes", testDiscountCodesCount)
	t.Run("Events", testEventsCount)
	t.Run("QuestionAnswers", testQuestionAnswersCount)
	t.Run("QuestionOptions", testQuestionOptionsCount)
	t.Run("Questions", testQuestionsCount)
	t.Run("TicketReservations", testTicketReservationsCount)
	t.Run("Tickets", testTicketsCount)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesCount)
	t.Run("TransactionItems", testTransactionItemsCount)
	t.Run("Transactions", testTransactionsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersHooks)
	t.Run("Accounts", testAccountsHooks)
	t.Run("Attendees", testAttendeesHooks)
	t.Run("Attributes", testAttributesHooks)
	t.Run("Customers", testCustomersHooks)
	t.Run("DiscountCodes", testDiscountCodesHooks)
	t.Run("Events", testEventsHooks)
	t.Run("QuestionAnswers", testQuestionAnswersHooks)
	t.Run("QuestionOptions", testQuestionOptionsHooks)
	t.Run("Questions", testQuestionsHooks)
	t.Run("TicketReservations", testTicketReservationsHooks)
	t.Run("Tickets", testTicketsHooks)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesHooks)
	t.Run("TransactionItems", testTransactionItemsHooks)
	t.Run("Transactions", testTransactionsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersInsert)
	t.Run("AccountUsers", testAccountUsersInsertWhitelist)
	t.Run("Accounts", testAccountsInsert)
	t.Run("Accounts", testAccountsInsertWhitelist)
	t.Run("Attendees", testAttendeesInsert)
	t.Run("Attendees", testAttendeesInsertWhitelist)
	t.Run("Attributes", testAttributesInsert)
	t.Run("Attributes", testAttributesInsertWhitelist)
	t.Run("Customers", testCustomersInsert)
	t.Run("Customers", testCustomersInsertWhitelist)
	t.Run("DiscountCodes", testDiscountCodesInsert)
	t.Run("DiscountCodes", testDiscountCodesInsertWhitelist)
	t.Run("Events", testEventsInsert)
	t.Run("Events", testEventsInsertWhitelist)
	t.Run("QuestionAnswers", testQuestionAnswersInsert)
	t.Run("QuestionAnswers", testQuestionAnswersInsertWhitelist)
	t.Run("QuestionOptions", testQuestionOptionsInsert)
	t.Run("QuestionOptions", testQuestionOptionsInsertWhitelist)
	t.Run("Questions", testQuestionsInsert)
	t.Run("Questions", testQuestionsInsertWhitelist)
	t.Run("TicketReservations", testTicketReservationsInsert)
	t.Run("TicketReservations", testTicketReservationsInsertWhitelist)
	t.Run("Tickets", testTicketsInsert)
	t.Run("Tickets", testTicketsInsertWhitelist)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesInsert)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesInsertWhitelist)
	t.Run("TransactionItems", testTransactionItemsInsert)
	t.Run("TransactionItems", testTransactionItemsInsertWhitelist)
	t.Run("Transactions", testTransactionsInsert)
	t.Run("Transactions", testTransactionsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AccountUserToAccountUsingAccount", testAccountUserToOneAccountUsingAccount)
	t.Run("AttendeeToTransactionUsingTransaction", testAttendeeToOneTransactionUsingTransaction)
	t.Run("AttendeeToTicketUsingTicket", testAttendeeToOneTicketUsingTicket)
	t.Run("AttendeeToEventUsingEvent", testAttendeeToOneEventUsingEvent)
	t.Run("CustomerToTransactionUsingTransaction", testCustomerToOneTransactionUsingTransaction)
	t.Run("CustomerToEventUsingEvent", testCustomerToOneEventUsingEvent)
	t.Run("EventToAccountUsingAccount", testEventToOneAccountUsingAccount)
	t.Run("EventToUserUsingUser", testEventToOneUserUsingUser)
	t.Run("QuestionAnswerToQuestionUsingQuestion", testQuestionAnswerToOneQuestionUsingQuestion)
	t.Run("QuestionOptionToQuestionUsingQuestion", testQuestionOptionToOneQuestionUsingQuestion)
	t.Run("TicketReservationToTicketUsingTicket", testTicketReservationToOneTicketUsingTicket)
	t.Run("TicketReservationToTransactionUsingTransaction", testTicketReservationToOneTransactionUsingTransaction)
	t.Run("TicketToEventUsingEvent", testTicketToOneEventUsingEvent)
	t.Run("TransactionDiscountCodeToTransactionUsingTransaction", testTransactionDiscountCodeToOneTransactionUsingTransaction)
	t.Run("TransactionDiscountCodeToDiscountCodeUsingDiscountCode", testTransactionDiscountCodeToOneDiscountCodeUsingDiscountCode)
	t.Run("TransactionItemToTransactionUsingTransaction", testTransactionItemToOneTransactionUsingTransaction)
	t.Run("TransactionItemToTicketUsingTicket", testTransactionItemToOneTicketUsingTicket)
	t.Run("TransactionToEventUsingEvent", testTransactionToOneEventUsingEvent)
	t.Run("UserToAccountUsingAccount", testUserToOneAccountUsingAccount)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AccountToAccountUsers", testAccountToManyAccountUsers)
	t.Run("AccountToEvents", testAccountToManyEvents)
	t.Run("AccountToUsers", testAccountToManyUsers)
	t.Run("AttributeToEvents", testAttributeToManyEvents)
	t.Run("AttributeToTickets", testAttributeToManyTickets)
	t.Run("AttributeToTransactions", testAttributeToManyTransactions)
	t.Run("DiscountCodeToTransactionDiscountCodes", testDiscountCodeToManyTransactionDiscountCodes)
	t.Run("EventToAttendees", testEventToManyAttendees)
	t.Run("EventToCustomers", testEventToManyCustomers)
	t.Run("EventToAttributes", testEventToManyAttributes)
	t.Run("EventToQuestions", testEventToManyQuestions)
	t.Run("EventToTickets", testEventToManyTickets)
	t.Run("EventToTransactions", testEventToManyTransactions)
	t.Run("QuestionToEvents", testQuestionToManyEvents)
	t.Run("QuestionToQuestionAnswers", testQuestionToManyQuestionAnswers)
	t.Run("QuestionToQuestionOptions", testQuestionToManyQuestionOptions)
	t.Run("QuestionToTickets", testQuestionToManyTickets)
	t.Run("TicketToAttendees", testTicketToManyAttendees)
	t.Run("TicketToAttributes", testTicketToManyAttributes)
	t.Run("TicketToQuestions", testTicketToManyQuestions)
	t.Run("TicketToTicketReservations", testTicketToManyTicketReservations)
	t.Run("TicketToTransactionItems", testTicketToManyTransactionItems)
	t.Run("TransactionToAttendees", testTransactionToManyAttendees)
	t.Run("TransactionToCustomers", testTransactionToManyCustomers)
	t.Run("TransactionToTicketReservations", testTransactionToManyTicketReservations)
	t.Run("TransactionToAttributes", testTransactionToManyAttributes)
	t.Run("TransactionToTransactionDiscountCodes", testTransactionToManyTransactionDiscountCodes)
	t.Run("TransactionToTransactionItems", testTransactionToManyTransactionItems)
	t.Run("UserToEvents", testUserToManyEvents)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AccountUserToAccountUsingAccountUsers", testAccountUserToOneSetOpAccountUsingAccount)
	t.Run("AttendeeToTransactionUsingAttendees", testAttendeeToOneSetOpTransactionUsingTransaction)
	t.Run("AttendeeToTicketUsingAttendees", testAttendeeToOneSetOpTicketUsingTicket)
	t.Run("AttendeeToEventUsingAttendees", testAttendeeToOneSetOpEventUsingEvent)
	t.Run("CustomerToTransactionUsingCustomers", testCustomerToOneSetOpTransactionUsingTransaction)
	t.Run("CustomerToEventUsingCustomers", testCustomerToOneSetOpEventUsingEvent)
	t.Run("EventToAccountUsingEvents", testEventToOneSetOpAccountUsingAccount)
	t.Run("EventToUserUsingEvents", testEventToOneSetOpUserUsingUser)
	t.Run("QuestionAnswerToQuestionUsingQuestionAnswers", testQuestionAnswerToOneSetOpQuestionUsingQuestion)
	t.Run("QuestionOptionToQuestionUsingQuestionOptions", testQuestionOptionToOneSetOpQuestionUsingQuestion)
	t.Run("TicketReservationToTicketUsingTicketReservations", testTicketReservationToOneSetOpTicketUsingTicket)
	t.Run("TicketReservationToTransactionUsingTicketReservations", testTicketReservationToOneSetOpTransactionUsingTransaction)
	t.Run("TicketToEventUsingTickets", testTicketToOneSetOpEventUsingEvent)
	t.Run("TransactionDiscountCodeToTransactionUsingTransactionDiscountCodes", testTransactionDiscountCodeToOneSetOpTransactionUsingTransaction)
	t.Run("TransactionDiscountCodeToDiscountCodeUsingTransactionDiscountCodes", testTransactionDiscountCodeToOneSetOpDiscountCodeUsingDiscountCode)
	t.Run("TransactionItemToTransactionUsingTransactionItems", testTransactionItemToOneSetOpTransactionUsingTransaction)
	t.Run("TransactionItemToTicketUsingTransactionItems", testTransactionItemToOneSetOpTicketUsingTicket)
	t.Run("TransactionToEventUsingTransactions", testTransactionToOneSetOpEventUsingEvent)
	t.Run("UserToAccountUsingUsers", testUserToOneSetOpAccountUsingAccount)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AccountToAccountUsers", testAccountToManyAddOpAccountUsers)
	t.Run("AccountToEvents", testAccountToManyAddOpEvents)
	t.Run("AccountToUsers", testAccountToManyAddOpUsers)
	t.Run("AttributeToEvents", testAttributeToManyAddOpEvents)
	t.Run("AttributeToTickets", testAttributeToManyAddOpTickets)
	t.Run("AttributeToTransactions", testAttributeToManyAddOpTransactions)
	t.Run("DiscountCodeToTransactionDiscountCodes", testDiscountCodeToManyAddOpTransactionDiscountCodes)
	t.Run("EventToAttendees", testEventToManyAddOpAttendees)
	t.Run("EventToCustomers", testEventToManyAddOpCustomers)
	t.Run("EventToAttributes", testEventToManyAddOpAttributes)
	t.Run("EventToQuestions", testEventToManyAddOpQuestions)
	t.Run("EventToTickets", testEventToManyAddOpTickets)
	t.Run("EventToTransactions", testEventToManyAddOpTransactions)
	t.Run("QuestionToEvents", testQuestionToManyAddOpEvents)
	t.Run("QuestionToQuestionAnswers", testQuestionToManyAddOpQuestionAnswers)
	t.Run("QuestionToQuestionOptions", testQuestionToManyAddOpQuestionOptions)
	t.Run("QuestionToTickets", testQuestionToManyAddOpTickets)
	t.Run("TicketToAttendees", testTicketToManyAddOpAttendees)
	t.Run("TicketToAttributes", testTicketToManyAddOpAttributes)
	t.Run("TicketToQuestions", testTicketToManyAddOpQuestions)
	t.Run("TicketToTicketReservations", testTicketToManyAddOpTicketReservations)
	t.Run("TicketToTransactionItems", testTicketToManyAddOpTransactionItems)
	t.Run("TransactionToAttendees", testTransactionToManyAddOpAttendees)
	t.Run("TransactionToCustomers", testTransactionToManyAddOpCustomers)
	t.Run("TransactionToTicketReservations", testTransactionToManyAddOpTicketReservations)
	t.Run("TransactionToAttributes", testTransactionToManyAddOpAttributes)
	t.Run("TransactionToTransactionDiscountCodes", testTransactionToManyAddOpTransactionDiscountCodes)
	t.Run("TransactionToTransactionItems", testTransactionToManyAddOpTransactionItems)
	t.Run("UserToEvents", testUserToManyAddOpEvents)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("AttributeToEvents", testAttributeToManySetOpEvents)
	t.Run("AttributeToTickets", testAttributeToManySetOpTickets)
	t.Run("AttributeToTransactions", testAttributeToManySetOpTransactions)
	t.Run("EventToAttributes", testEventToManySetOpAttributes)
	t.Run("EventToQuestions", testEventToManySetOpQuestions)
	t.Run("QuestionToEvents", testQuestionToManySetOpEvents)
	t.Run("QuestionToTickets", testQuestionToManySetOpTickets)
	t.Run("TicketToAttributes", testTicketToManySetOpAttributes)
	t.Run("TicketToQuestions", testTicketToManySetOpQuestions)
	t.Run("TransactionToAttributes", testTransactionToManySetOpAttributes)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("AttributeToEvents", testAttributeToManyRemoveOpEvents)
	t.Run("AttributeToTickets", testAttributeToManyRemoveOpTickets)
	t.Run("AttributeToTransactions", testAttributeToManyRemoveOpTransactions)
	t.Run("EventToAttributes", testEventToManyRemoveOpAttributes)
	t.Run("EventToQuestions", testEventToManyRemoveOpQuestions)
	t.Run("QuestionToEvents", testQuestionToManyRemoveOpEvents)
	t.Run("QuestionToTickets", testQuestionToManyRemoveOpTickets)
	t.Run("TicketToAttributes", testTicketToManyRemoveOpAttributes)
	t.Run("TicketToQuestions", testTicketToManyRemoveOpQuestions)
	t.Run("TransactionToAttributes", testTransactionToManyRemoveOpAttributes)
}

func TestReload(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersReload)
	t.Run("Accounts", testAccountsReload)
	t.Run("Attendees", testAttendeesReload)
	t.Run("Attributes", testAttributesReload)
	t.Run("Customers", testCustomersReload)
	t.Run("DiscountCodes", testDiscountCodesReload)
	t.Run("Events", testEventsReload)
	t.Run("QuestionAnswers", testQuestionAnswersReload)
	t.Run("QuestionOptions", testQuestionOptionsReload)
	t.Run("Questions", testQuestionsReload)
	t.Run("TicketReservations", testTicketReservationsReload)
	t.Run("Tickets", testTicketsReload)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesReload)
	t.Run("TransactionItems", testTransactionItemsReload)
	t.Run("Transactions", testTransactionsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersReloadAll)
	t.Run("Accounts", testAccountsReloadAll)
	t.Run("Attendees", testAttendeesReloadAll)
	t.Run("Attributes", testAttributesReloadAll)
	t.Run("Customers", testCustomersReloadAll)
	t.Run("DiscountCodes", testDiscountCodesReloadAll)
	t.Run("Events", testEventsReloadAll)
	t.Run("QuestionAnswers", testQuestionAnswersReloadAll)
	t.Run("QuestionOptions", testQuestionOptionsReloadAll)
	t.Run("Questions", testQuestionsReloadAll)
	t.Run("TicketReservations", testTicketReservationsReloadAll)
	t.Run("Tickets", testTicketsReloadAll)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesReloadAll)
	t.Run("TransactionItems", testTransactionItemsReloadAll)
	t.Run("Transactions", testTransactionsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersSelect)
	t.Run("Accounts", testAccountsSelect)
	t.Run("Attendees", testAttendeesSelect)
	t.Run("Attributes", testAttributesSelect)
	t.Run("Customers", testCustomersSelect)
	t.Run("DiscountCodes", testDiscountCodesSelect)
	t.Run("Events", testEventsSelect)
	t.Run("QuestionAnswers", testQuestionAnswersSelect)
	t.Run("QuestionOptions", testQuestionOptionsSelect)
	t.Run("Questions", testQuestionsSelect)
	t.Run("TicketReservations", testTicketReservationsSelect)
	t.Run("Tickets", testTicketsSelect)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesSelect)
	t.Run("TransactionItems", testTransactionItemsSelect)
	t.Run("Transactions", testTransactionsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersUpdate)
	t.Run("Accounts", testAccountsUpdate)
	t.Run("Attendees", testAttendeesUpdate)
	t.Run("Attributes", testAttributesUpdate)
	t.Run("Customers", testCustomersUpdate)
	t.Run("DiscountCodes", testDiscountCodesUpdate)
	t.Run("Events", testEventsUpdate)
	t.Run("QuestionAnswers", testQuestionAnswersUpdate)
	t.Run("QuestionOptions", testQuestionOptionsUpdate)
	t.Run("Questions", testQuestionsUpdate)
	t.Run("TicketReservations", testTicketReservationsUpdate)
	t.Run("Tickets", testTicketsUpdate)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesUpdate)
	t.Run("TransactionItems", testTransactionItemsUpdate)
	t.Run("Transactions", testTransactionsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersSliceUpdateAll)
	t.Run("Accounts", testAccountsSliceUpdateAll)
	t.Run("Attendees", testAttendeesSliceUpdateAll)
	t.Run("Attributes", testAttributesSliceUpdateAll)
	t.Run("Customers", testCustomersSliceUpdateAll)
	t.Run("DiscountCodes", testDiscountCodesSliceUpdateAll)
	t.Run("Events", testEventsSliceUpdateAll)
	t.Run("QuestionAnswers", testQuestionAnswersSliceUpdateAll)
	t.Run("QuestionOptions", testQuestionOptionsSliceUpdateAll)
	t.Run("Questions", testQuestionsSliceUpdateAll)
	t.Run("TicketReservations", testTicketReservationsSliceUpdateAll)
	t.Run("Tickets", testTicketsSliceUpdateAll)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesSliceUpdateAll)
	t.Run("TransactionItems", testTransactionItemsSliceUpdateAll)
	t.Run("Transactions", testTransactionsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
