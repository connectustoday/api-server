import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';
import INotification from './notification';
import IAccountSettings from './account-settings';

export default class IAccount extends Typegoose {

	@prop({ required: true })
	schema_version: number;

	@prop({ required: true })
	id: string;

	@prop({ required: true })
	username: string;

	@prop({ required: true })
	email: string;

	@prop({ required: true })
	password: string;

	@prop()
	oauth_token?: string;

	@prop()
	oauth_service?: string;

	@prop({ required: true })
	is_email_verified: boolean;

	@prop({ required: true })
	last_login: number;

	@prop({ required: true })
	notifications: Array<INotification>;

	@prop({ required: true })
	avatar: string;

	@prop({ required: true })
	header: string;

	@prop({ required: true })
	created_at: number;

	@prop()
	pending_connections?: Array<string>;

	@prop()
	requested_connections?: Array<string>;

	@prop()
	posts?: Array<string>;

	@prop()
	liked?: Array<[string, number]>;

	@prop()
	shared?: Array<[string, number]>;

	@prop({ required: true })
	settings: IAccountSettings;

	@prop()
	admin_note?: string;

}
