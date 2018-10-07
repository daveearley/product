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
	t.Run("EventAttributes", testEventAttributes)
	t.Run("EventQuestions", testEventQuestions)
	t.Run("Events", testEvents)
	t.Run("QuestionAnswers", testQuestionAnswers)
	t.Run("Questions", testQuestions)
	t.Run("TicketAttributes", testTicketAttributes)
	t.Run("TicketQuestions", testTicketQuestions)
	t.Run("Tickets", testTickets)
	t.Run("TransactionAttributes", testTransactionAttributes)
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
	t.Run("EventAttributes", testEventAttributesDelete)
	t.Run("EventQuestions", testEventQuestionsDelete)
	t.Run("Events", testEventsDelete)
	t.Run("QuestionAnswers", testQuestionAnswersDelete)
	t.Run("Questions", testQuestionsDelete)
	t.Run("TicketAttributes", testTicketAttributesDelete)
	t.Run("TicketQuestions", testTicketQuestionsDelete)
	t.Run("Tickets", testTicketsDelete)
	t.Run("TransactionAttributes", testTransactionAttributesDelete)
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
	t.Run("EventAttributes", testEventAttributesQueryDeleteAll)
	t.Run("EventQuestions", testEventQuestionsQueryDeleteAll)
	t.Run("Events", testEventsQueryDeleteAll)
	t.Run("QuestionAnswers", testQuestionAnswersQueryDeleteAll)
	t.Run("Questions", testQuestionsQueryDeleteAll)
	t.Run("TicketAttributes", testTicketAttributesQueryDeleteAll)
	t.Run("TicketQuestions", testTicketQuestionsQueryDeleteAll)
	t.Run("Tickets", testTicketsQueryDeleteAll)
	t.Run("TransactionAttributes", testTransactionAttributesQueryDeleteAll)
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
	t.Run("EventAttributes", testEventAttributesSliceDeleteAll)
	t.Run("EventQuestions", testEventQuestionsSliceDeleteAll)
	t.Run("Events", testEventsSliceDeleteAll)
	t.Run("QuestionAnswers", testQuestionAnswersSliceDeleteAll)
	t.Run("Questions", testQuestionsSliceDeleteAll)
	t.Run("TicketAttributes", testTicketAttributesSliceDeleteAll)
	t.Run("TicketQuestions", testTicketQuestionsSliceDeleteAll)
	t.Run("Tickets", testTicketsSliceDeleteAll)
	t.Run("TransactionAttributes", testTransactionAttributesSliceDeleteAll)
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
	t.Run("EventAttributes", testEventAttributesExists)
	t.Run("EventQuestions", testEventQuestionsExists)
	t.Run("Events", testEventsExists)
	t.Run("QuestionAnswers", testQuestionAnswersExists)
	t.Run("Questions", testQuestionsExists)
	t.Run("TicketAttributes", testTicketAttributesExists)
	t.Run("TicketQuestions", testTicketQuestionsExists)
	t.Run("Tickets", testTicketsExists)
	t.Run("TransactionAttributes", testTransactionAttributesExists)
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
	t.Run("EventAttributes", testEventAttributesFind)
	t.Run("EventQuestions", testEventQuestionsFind)
	t.Run("Events", testEventsFind)
	t.Run("QuestionAnswers", testQuestionAnswersFind)
	t.Run("Questions", testQuestionsFind)
	t.Run("TicketAttributes", testTicketAttributesFind)
	t.Run("TicketQuestions", testTicketQuestionsFind)
	t.Run("Tickets", testTicketsFind)
	t.Run("TransactionAttributes", testTransactionAttributesFind)
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
	t.Run("EventAttributes", testEventAttributesBind)
	t.Run("EventQuestions", testEventQuestionsBind)
	t.Run("Events", testEventsBind)
	t.Run("QuestionAnswers", testQuestionAnswersBind)
	t.Run("Questions", testQuestionsBind)
	t.Run("TicketAttributes", testTicketAttributesBind)
	t.Run("TicketQuestions", testTicketQuestionsBind)
	t.Run("Tickets", testTicketsBind)
	t.Run("TransactionAttributes", testTransactionAttributesBind)
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
	t.Run("EventAttributes", testEventAttributesOne)
	t.Run("EventQuestions", testEventQuestionsOne)
	t.Run("Events", testEventsOne)
	t.Run("QuestionAnswers", testQuestionAnswersOne)
	t.Run("Questions", testQuestionsOne)
	t.Run("TicketAttributes", testTicketAttributesOne)
	t.Run("TicketQuestions", testTicketQuestionsOne)
	t.Run("Tickets", testTicketsOne)
	t.Run("TransactionAttributes", testTransactionAttributesOne)
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
	t.Run("EventAttributes", testEventAttributesAll)
	t.Run("EventQuestions", testEventQuestionsAll)
	t.Run("Events", testEventsAll)
	t.Run("QuestionAnswers", testQuestionAnswersAll)
	t.Run("Questions", testQuestionsAll)
	t.Run("TicketAttributes", testTicketAttributesAll)
	t.Run("TicketQuestions", testTicketQuestionsAll)
	t.Run("Tickets", testTicketsAll)
	t.Run("TransactionAttributes", testTransactionAttributesAll)
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
	t.Run("EventAttributes", testEventAttributesCount)
	t.Run("EventQuestions", testEventQuestionsCount)
	t.Run("Events", testEventsCount)
	t.Run("QuestionAnswers", testQuestionAnswersCount)
	t.Run("Questions", testQuestionsCount)
	t.Run("TicketAttributes", testTicketAttributesCount)
	t.Run("TicketQuestions", testTicketQuestionsCount)
	t.Run("Tickets", testTicketsCount)
	t.Run("TransactionAttributes", testTransactionAttributesCount)
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
	t.Run("EventAttributes", testEventAttributesHooks)
	t.Run("EventQuestions", testEventQuestionsHooks)
	t.Run("Events", testEventsHooks)
	t.Run("QuestionAnswers", testQuestionAnswersHooks)
	t.Run("Questions", testQuestionsHooks)
	t.Run("TicketAttributes", testTicketAttributesHooks)
	t.Run("TicketQuestions", testTicketQuestionsHooks)
	t.Run("Tickets", testTicketsHooks)
	t.Run("TransactionAttributes", testTransactionAttributesHooks)
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
	t.Run("EventAttributes", testEventAttributesInsert)
	t.Run("EventAttributes", testEventAttributesInsertWhitelist)
	t.Run("EventQuestions", testEventQuestionsInsert)
	t.Run("EventQuestions", testEventQuestionsInsertWhitelist)
	t.Run("Events", testEventsInsert)
	t.Run("Events", testEventsInsertWhitelist)
	t.Run("QuestionAnswers", testQuestionAnswersInsert)
	t.Run("QuestionAnswers", testQuestionAnswersInsertWhitelist)
	t.Run("Questions", testQuestionsInsert)
	t.Run("Questions", testQuestionsInsertWhitelist)
	t.Run("TicketAttributes", testTicketAttributesInsert)
	t.Run("TicketAttributes", testTicketAttributesInsertWhitelist)
	t.Run("TicketQuestions", testTicketQuestionsInsert)
	t.Run("TicketQuestions", testTicketQuestionsInsertWhitelist)
	t.Run("Tickets", testTicketsInsert)
	t.Run("Tickets", testTicketsInsertWhitelist)
	t.Run("TransactionAttributes", testTransactionAttributesInsert)
	t.Run("TransactionAttributes", testTransactionAttributesInsertWhitelist)
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
	t.Run("AttendeeToCustomerUsingCustomer", testAttendeeToOneCustomerUsingCustomer)
	t.Run("AttendeeToTicketUsingTicket", testAttendeeToOneTicketUsingTicket)
	t.Run("AttendeeToEventUsingEvent", testAttendeeToOneEventUsingEvent)
	t.Run("CustomerToTransactionUsingTransaction", testCustomerToOneTransactionUsingTransaction)
	t.Run("CustomerToEventUsingEvent", testCustomerToOneEventUsingEvent)
	t.Run("EventAttributeToEventUsingEvent", testEventAttributeToOneEventUsingEvent)
	t.Run("EventAttributeToAttributeUsingAttribute", testEventAttributeToOneAttributeUsingAttribute)
	t.Run("EventQuestionToEventUsingEvent", testEventQuestionToOneEventUsingEvent)
	t.Run("EventQuestionToQuestionUsingQuestion", testEventQuestionToOneQuestionUsingQuestion)
	t.Run("EventToAccountUserUsingUser", testEventToOneAccountUserUsingUser)
	t.Run("EventToAccountUsingAccount", testEventToOneAccountUsingAccount)
	t.Run("QuestionAnswerToQuestionUsingQuestion", testQuestionAnswerToOneQuestionUsingQuestion)
	t.Run("TicketAttributeToTicketUsingTicket", testTicketAttributeToOneTicketUsingTicket)
	t.Run("TicketAttributeToAttributeUsingAttribute", testTicketAttributeToOneAttributeUsingAttribute)
	t.Run("TicketQuestionToTicketUsingTicket", testTicketQuestionToOneTicketUsingTicket)
	t.Run("TicketQuestionToQuestionUsingQuestion", testTicketQuestionToOneQuestionUsingQuestion)
	t.Run("TicketToEventUsingEvent", testTicketToOneEventUsingEvent)
	t.Run("TransactionAttributeToTransactionUsingTransaction", testTransactionAttributeToOneTransactionUsingTransaction)
	t.Run("TransactionAttributeToAttributeUsingAttribute", testTransactionAttributeToOneAttributeUsingAttribute)
	t.Run("TransactionDiscountCodeToTransactionUsingTransaction", testTransactionDiscountCodeToOneTransactionUsingTransaction)
	t.Run("TransactionDiscountCodeToDiscountCodeUsingDiscountCode", testTransactionDiscountCodeToOneDiscountCodeUsingDiscountCode)
	t.Run("TransactionItemToTransactionUsingTransaction", testTransactionItemToOneTransactionUsingTransaction)
	t.Run("TransactionItemToTicketUsingTicket", testTransactionItemToOneTicketUsingTicket)
	t.Run("TransactionToEventUsingEvent", testTransactionToOneEventUsingEvent)
	t.Run("TransactionToCustomerUsingCustomer", testTransactionToOneCustomerUsingCustomer)
	t.Run("UserToAccountUsingAccount", testUserToOneAccountUsingAccount)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AccountUserToUserEvents", testAccountUserToManyUserEvents)
	t.Run("AccountToAccountUsers", testAccountToManyAccountUsers)
	t.Run("AccountToEvents", testAccountToManyEvents)
	t.Run("AccountToUsers", testAccountToManyUsers)
	t.Run("AttributeToEventAttributes", testAttributeToManyEventAttributes)
	t.Run("AttributeToTicketAttributes", testAttributeToManyTicketAttributes)
	t.Run("AttributeToTransactionAttributes", testAttributeToManyTransactionAttributes)
	t.Run("CustomerToAttendees", testCustomerToManyAttendees)
	t.Run("CustomerToTransactions", testCustomerToManyTransactions)
	t.Run("DiscountCodeToTransactionDiscountCodes", testDiscountCodeToManyTransactionDiscountCodes)
	t.Run("EventToAttendees", testEventToManyAttendees)
	t.Run("EventToCustomers", testEventToManyCustomers)
	t.Run("EventToEventAttributes", testEventToManyEventAttributes)
	t.Run("EventToEventQuestions", testEventToManyEventQuestions)
	t.Run("EventToTickets", testEventToManyTickets)
	t.Run("EventToTransactions", testEventToManyTransactions)
	t.Run("QuestionToEventQuestions", testQuestionToManyEventQuestions)
	t.Run("QuestionToQuestionAnswers", testQuestionToManyQuestionAnswers)
	t.Run("QuestionToTicketQuestions", testQuestionToManyTicketQuestions)
	t.Run("TicketToAttendees", testTicketToManyAttendees)
	t.Run("TicketToTicketAttributes", testTicketToManyTicketAttributes)
	t.Run("TicketToTicketQuestions", testTicketToManyTicketQuestions)
	t.Run("TicketToTransactionItems", testTicketToManyTransactionItems)
	t.Run("TransactionToCustomers", testTransactionToManyCustomers)
	t.Run("TransactionToTransactionAttributes", testTransactionToManyTransactionAttributes)
	t.Run("TransactionToTransactionDiscountCodes", testTransactionToManyTransactionDiscountCodes)
	t.Run("TransactionToTransactionItems", testTransactionToManyTransactionItems)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AccountUserToAccountUsingAccountUsers", testAccountUserToOneSetOpAccountUsingAccount)
	t.Run("AttendeeToCustomerUsingAttendees", testAttendeeToOneSetOpCustomerUsingCustomer)
	t.Run("AttendeeToTicketUsingAttendees", testAttendeeToOneSetOpTicketUsingTicket)
	t.Run("AttendeeToEventUsingAttendees", testAttendeeToOneSetOpEventUsingEvent)
	t.Run("CustomerToTransactionUsingCustomers", testCustomerToOneSetOpTransactionUsingTransaction)
	t.Run("CustomerToEventUsingCustomers", testCustomerToOneSetOpEventUsingEvent)
	t.Run("EventAttributeToEventUsingEventAttributes", testEventAttributeToOneSetOpEventUsingEvent)
	t.Run("EventAttributeToAttributeUsingEventAttributes", testEventAttributeToOneSetOpAttributeUsingAttribute)
	t.Run("EventQuestionToEventUsingEventQuestions", testEventQuestionToOneSetOpEventUsingEvent)
	t.Run("EventQuestionToQuestionUsingEventQuestions", testEventQuestionToOneSetOpQuestionUsingQuestion)
	t.Run("EventToAccountUserUsingUserEvents", testEventToOneSetOpAccountUserUsingUser)
	t.Run("EventToAccountUsingEvents", testEventToOneSetOpAccountUsingAccount)
	t.Run("QuestionAnswerToQuestionUsingQuestionAnswers", testQuestionAnswerToOneSetOpQuestionUsingQuestion)
	t.Run("TicketAttributeToTicketUsingTicketAttributes", testTicketAttributeToOneSetOpTicketUsingTicket)
	t.Run("TicketAttributeToAttributeUsingTicketAttributes", testTicketAttributeToOneSetOpAttributeUsingAttribute)
	t.Run("TicketQuestionToTicketUsingTicketQuestions", testTicketQuestionToOneSetOpTicketUsingTicket)
	t.Run("TicketQuestionToQuestionUsingTicketQuestions", testTicketQuestionToOneSetOpQuestionUsingQuestion)
	t.Run("TicketToEventUsingTickets", testTicketToOneSetOpEventUsingEvent)
	t.Run("TransactionAttributeToTransactionUsingTransactionAttributes", testTransactionAttributeToOneSetOpTransactionUsingTransaction)
	t.Run("TransactionAttributeToAttributeUsingTransactionAttributes", testTransactionAttributeToOneSetOpAttributeUsingAttribute)
	t.Run("TransactionDiscountCodeToTransactionUsingTransactionDiscountCodes", testTransactionDiscountCodeToOneSetOpTransactionUsingTransaction)
	t.Run("TransactionDiscountCodeToDiscountCodeUsingTransactionDiscountCodes", testTransactionDiscountCodeToOneSetOpDiscountCodeUsingDiscountCode)
	t.Run("TransactionItemToTransactionUsingTransactionItems", testTransactionItemToOneSetOpTransactionUsingTransaction)
	t.Run("TransactionItemToTicketUsingTransactionItems", testTransactionItemToOneSetOpTicketUsingTicket)
	t.Run("TransactionToEventUsingTransactions", testTransactionToOneSetOpEventUsingEvent)
	t.Run("TransactionToCustomerUsingTransactions", testTransactionToOneSetOpCustomerUsingCustomer)
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
	t.Run("AccountUserToUserEvents", testAccountUserToManyAddOpUserEvents)
	t.Run("AccountToAccountUsers", testAccountToManyAddOpAccountUsers)
	t.Run("AccountToEvents", testAccountToManyAddOpEvents)
	t.Run("AccountToUsers", testAccountToManyAddOpUsers)
	t.Run("AttributeToEventAttributes", testAttributeToManyAddOpEventAttributes)
	t.Run("AttributeToTicketAttributes", testAttributeToManyAddOpTicketAttributes)
	t.Run("AttributeToTransactionAttributes", testAttributeToManyAddOpTransactionAttributes)
	t.Run("CustomerToAttendees", testCustomerToManyAddOpAttendees)
	t.Run("CustomerToTransactions", testCustomerToManyAddOpTransactions)
	t.Run("DiscountCodeToTransactionDiscountCodes", testDiscountCodeToManyAddOpTransactionDiscountCodes)
	t.Run("EventToAttendees", testEventToManyAddOpAttendees)
	t.Run("EventToCustomers", testEventToManyAddOpCustomers)
	t.Run("EventToEventAttributes", testEventToManyAddOpEventAttributes)
	t.Run("EventToEventQuestions", testEventToManyAddOpEventQuestions)
	t.Run("EventToTickets", testEventToManyAddOpTickets)
	t.Run("EventToTransactions", testEventToManyAddOpTransactions)
	t.Run("QuestionToEventQuestions", testQuestionToManyAddOpEventQuestions)
	t.Run("QuestionToQuestionAnswers", testQuestionToManyAddOpQuestionAnswers)
	t.Run("QuestionToTicketQuestions", testQuestionToManyAddOpTicketQuestions)
	t.Run("TicketToAttendees", testTicketToManyAddOpAttendees)
	t.Run("TicketToTicketAttributes", testTicketToManyAddOpTicketAttributes)
	t.Run("TicketToTicketQuestions", testTicketToManyAddOpTicketQuestions)
	t.Run("TicketToTransactionItems", testTicketToManyAddOpTransactionItems)
	t.Run("TransactionToCustomers", testTransactionToManyAddOpCustomers)
	t.Run("TransactionToTransactionAttributes", testTransactionToManyAddOpTransactionAttributes)
	t.Run("TransactionToTransactionDiscountCodes", testTransactionToManyAddOpTransactionDiscountCodes)
	t.Run("TransactionToTransactionItems", testTransactionToManyAddOpTransactionItems)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("AccountUsers", testAccountUsersReload)
	t.Run("Accounts", testAccountsReload)
	t.Run("Attendees", testAttendeesReload)
	t.Run("Attributes", testAttributesReload)
	t.Run("Customers", testCustomersReload)
	t.Run("DiscountCodes", testDiscountCodesReload)
	t.Run("EventAttributes", testEventAttributesReload)
	t.Run("EventQuestions", testEventQuestionsReload)
	t.Run("Events", testEventsReload)
	t.Run("QuestionAnswers", testQuestionAnswersReload)
	t.Run("Questions", testQuestionsReload)
	t.Run("TicketAttributes", testTicketAttributesReload)
	t.Run("TicketQuestions", testTicketQuestionsReload)
	t.Run("Tickets", testTicketsReload)
	t.Run("TransactionAttributes", testTransactionAttributesReload)
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
	t.Run("EventAttributes", testEventAttributesReloadAll)
	t.Run("EventQuestions", testEventQuestionsReloadAll)
	t.Run("Events", testEventsReloadAll)
	t.Run("QuestionAnswers", testQuestionAnswersReloadAll)
	t.Run("Questions", testQuestionsReloadAll)
	t.Run("TicketAttributes", testTicketAttributesReloadAll)
	t.Run("TicketQuestions", testTicketQuestionsReloadAll)
	t.Run("Tickets", testTicketsReloadAll)
	t.Run("TransactionAttributes", testTransactionAttributesReloadAll)
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
	t.Run("EventAttributes", testEventAttributesSelect)
	t.Run("EventQuestions", testEventQuestionsSelect)
	t.Run("Events", testEventsSelect)
	t.Run("QuestionAnswers", testQuestionAnswersSelect)
	t.Run("Questions", testQuestionsSelect)
	t.Run("TicketAttributes", testTicketAttributesSelect)
	t.Run("TicketQuestions", testTicketQuestionsSelect)
	t.Run("Tickets", testTicketsSelect)
	t.Run("TransactionAttributes", testTransactionAttributesSelect)
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
	t.Run("EventAttributes", testEventAttributesUpdate)
	t.Run("EventQuestions", testEventQuestionsUpdate)
	t.Run("Events", testEventsUpdate)
	t.Run("QuestionAnswers", testQuestionAnswersUpdate)
	t.Run("Questions", testQuestionsUpdate)
	t.Run("TicketAttributes", testTicketAttributesUpdate)
	t.Run("TicketQuestions", testTicketQuestionsUpdate)
	t.Run("Tickets", testTicketsUpdate)
	t.Run("TransactionAttributes", testTransactionAttributesUpdate)
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
	t.Run("EventAttributes", testEventAttributesSliceUpdateAll)
	t.Run("EventQuestions", testEventQuestionsSliceUpdateAll)
	t.Run("Events", testEventsSliceUpdateAll)
	t.Run("QuestionAnswers", testQuestionAnswersSliceUpdateAll)
	t.Run("Questions", testQuestionsSliceUpdateAll)
	t.Run("TicketAttributes", testTicketAttributesSliceUpdateAll)
	t.Run("TicketQuestions", testTicketQuestionsSliceUpdateAll)
	t.Run("Tickets", testTicketsSliceUpdateAll)
	t.Run("TransactionAttributes", testTransactionAttributesSliceUpdateAll)
	t.Run("TransactionDiscountCodes", testTransactionDiscountCodesSliceUpdateAll)
	t.Run("TransactionItems", testTransactionItemsSliceUpdateAll)
	t.Run("Transactions", testTransactionsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}